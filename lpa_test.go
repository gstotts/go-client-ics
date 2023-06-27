package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLPA_ListPrincipalActivity(t *testing.T) {
	testCases := []struct {
		test_name              string
		principal_id           string
		start                  string
		end                    string
		resp_file              string
		executed_actions_count int
		prinicpal_name         string
		resource_type          string
		err_expected           bool
	}{
		{"Valid Request", "serviceuser:12:FFFFFFFFFFF:", "2022-08-14", "2022-08-16", "lpa/list_principal_activity_response.json", 3, "ics-assume-role-test", "serviceuser", false},
		{"Invalid Start Format", "serviceuser:4:FFFFFFFFFFF:", "20230109", "20223-01-10", "", 0, "", "", true},
		{"Invalid End Format", "serviceuser:4:FFFFFFFFFFF:", "2023-01-09", "2023/1/10", "", 0, "", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			setup()
			mux.HandleFunc(fmt.Sprintf("/v3/lpa/principals/%s/actions", tc.principal_id), func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tc.start, r.URL.Query().Get("start"))
				assert.Equal(t, tc.end, r.URL.Query().Get("end"))
				assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, getJSONFile(tc.resp_file))
			})

			activity, err := client.ListPrincipalActivity(tc.principal_id, tc.start, tc.end)
			if tc.err_expected {
				fmt.Println(err)
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.executed_actions_count, len(activity.ExecutedActions))
				assert.Equal(t, tc.resource_type, activity.Principal.ResourceType)
				assert.Equal(t, tc.prinicpal_name, activity.Principal.Name)
				assert.Equal(t, tc.principal_id, activity.Principal.ResourceID)
			}
			teardown()
		})
	}
}

func TestLPA_ListPrincipalPermissions(t *testing.T) {
	testCases := []struct {
		test_name    string
		principal_id string
		err_expected bool
	}{
		{"Valid_Request", "serviceuser:4:FFFFFFFFFFF:", false},
		{"Invalid Principal Request", "servicedog1:123:", true},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			setup()
			mux.HandleFunc("/v3/lpa/principals/serviceuser:4:FFFFFFFFFFF:/permissions", func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, getJSONFile("lpa/list_principal_permissions_response.json"))
			})

			result, err := client.ListPrincipalPermissions(tc.principal_id)
			if tc.err_expected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "2023-06-26", result.End)
			}
			teardown()
		})
	}
}

func TestLPA_GenerateDenyNotActionPolicy(t *testing.T) {
	testCases := []struct {
		test_name    string
		principal_id string
		start        string
		end          string
		resp_file    string
		sid          string
		notactions   []string
		resources    []string
		err_expected bool
	}{
		{"Valid Request", "serviceuser:12:FFFFFFFFFFF:", "2022-08-14", "2022-08-16", "lpa/generate_notaction_response.json", "LPADenyUnusedPermissions", []string{"s3:GetObject", "s3:Delete*"}, []string{"*"}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			setup()
			mux.HandleFunc(fmt.Sprintf("/v3/lpa/principals/%s/access-remediate-policy", tc.principal_id), func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tc.start, r.URL.Query().Get("start"))
				assert.Equal(t, tc.end, r.URL.Query().Get("end"))
				assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, getJSONFile(tc.resp_file))
			})

			policy, err := client.GenerateDenyNotActionPolicy(tc.principal_id, tc.start, tc.end)
			if tc.err_expected {
				fmt.Println(err)
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "Deny", policy.Policy.Statement[0].Effect)
				assert.Equal(t, tc.sid, policy.Policy.Statement[0].Sid)
				assert.Equal(t, tc.notactions, policy.Policy.Statement[0].NotAction)
				assert.Equal(t, tc.resources, policy.Policy.Statement[0].Resource)
				assert.Equal(t, "2012-10-17", policy.Policy.Version)
			}
			teardown()
		})
	}
}
