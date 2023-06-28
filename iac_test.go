package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIAC_ListIACConfigs(t *testing.T) {
	testCases := []struct {
		id          int
		orgId       int
		name        string
		description string
		packid      int
		source      string
		blacklist   []string
		warn        []string
		slack       string
		email       []string
		lastBuild   string
		created     string
		modified    string
		total       int
		success     int
		fail        int
		url         string
		favorite    bool
		tfcCount    int
		exceptions  bool
	}{
		{1, 1, "Initial Testing - AWS Foundational", "Initial IaC Scan of AWS Foundational Security Best Practices", 45, "backoffice", []string{"backoffice:27", "backoffice:7"}, []string{"backoffice:55", "backoffice:72"}, "", []string{}, "", "2023-06-28T03:07:37Z", "2023-06-28T03:07:37Z", 0, 0, 0, "http://localhost:8001/v3/iac/scan", false, 0, true},
	}

	for i, tc := range testCases {
		t.Run("Valid Results", func(t *testing.T) {
			setup()
			mux.HandleFunc("/v3/iac/configs", func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, getJSONFile("iac/sample_list_configs.json"))
			})
			resp, err := client.ListIACConfigs()
			assert.NoError(t, err)
			assert.Equal(t, tc.id, resp[i].ID)
			assert.Equal(t, tc.orgId, resp[i].OrganizationID)
			assert.Equal(t, tc.name, resp[i].Name)
			assert.Equal(t, tc.description, resp[i].Description)
			assert.Equal(t, tc.packid, resp[i].PackID)
			assert.Equal(t, tc.source, resp[i].Source)
			assert.Equal(t, tc.blacklist, resp[i].InsightsBlacklist)
			assert.Equal(t, tc.warn, resp[i].InsightsWarnOnly)
			assert.Equal(t, tc.slack, resp[i].SlackChannel)
			assert.Equal(t, tc.email, resp[i].EmailRecipients)
			assert.Equal(t, tc.lastBuild, resp[i].LastBuild)
			assert.Equal(t, tc.created, resp[i].Created)
			assert.Equal(t, tc.modified, resp[i].Modified)
			assert.Equal(t, tc.total, resp[i].TotalBuilds)
			assert.Equal(t, tc.success, resp[i].SuccessCount)
			assert.Equal(t, tc.fail, resp[i].FailureCount)
			assert.Equal(t, tc.url, resp[i].ConsumptionURL)
			assert.Equal(t, tc.favorite, resp[i].Favorite)
			assert.Equal(t, tc.tfcCount, resp[i].TFCCount)
			assert.Equal(t, tc.exceptions, resp[i].LocalExceptionsEnabled)
			teardown()
		})
	}
}
