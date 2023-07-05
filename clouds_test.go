package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClouds_ListCloudTypes(t *testing.T) {
	testCases := []struct {
		cloud_type_name string
		cloud_type_id   string
		cloud_access    string
	}{
		{"Alibaba Cloud", "ALICLOUD", "public"},
		{"Amazon Web Services", "AWS", "public"},
		{"Amazon Web Services (China)", "AWS_CHINA", "public"},
		{"Amazon Web Services (GovCloud)", "AWS_GOV", "public"},
		{"Microsoft Azure", "AZURE_ARM", "public"},
		{"Microsoft Azure (China)", "AZURE_CHINA", "public"},
		{"Microsoft Azure (GovCloud)", "AZURE_GOV", "public"},
		{"Google Cloud Platform", "GCE", "public"},
		{"Kubernetes Security", "K8S_R7", "public"},
		{"Oracle Cloud", "OCI", "public"},
	}

	setup()
	mux.HandleFunc("/v2/public/cloudtypes/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/list_cloudtypes.json"))
	})
	resp, err := client.ListCloudTypes()
	assert.NoError(t, err)
	assert.NotEqual(t, len(resp.Clouds), 0)

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Validating_%s", tc.cloud_type_id), func(t *testing.T) {
			assert.Equal(t, tc.cloud_type_id, resp.Clouds[i].ID)
			assert.Equal(t, tc.cloud_type_name, resp.Clouds[i].Name)
			assert.Equal(t, tc.cloud_access, resp.Clouds[i].Access)
		})
	}
	teardown()
}

func TestClouds_ListClouds(t *testing.T) {
	testCases := []struct {
		id                  int
		name                string
		type_id             string
		account_id          string
		create              string
		host_assess_status  []string
		host_assess_enabled bool
		status              string
		badges              int
		resources           int
		refreshed           string
		role                string
		group_resource_id   string
		resource_id         string
		edh_role            string
		strategy_id         int
		org_id              string
		org_dn              string
		org_nickname        string
	}{
		{1, "My GCP", "GCE", "gcp", "2023-06-09 05:21:06", []string{}, false, "REFRESH", 1, 231, "2023-07-05 15:59:51", "insightcloudsec@gcp.iam.gserviceaccount.com", "divvyorganizationservice:1", "divvyorganizationservice:1", "idle", 3, "", "", ""},
		{3, "AWS Root", "AWS", "123456789101", "2023-06-26 23:09:20", []string{}, false, "DEFAULT", 1, 233, "2023-07-05 15:59:51", "insightcloudsec_sts_assume_role", "divvyorganizationservice:3", "divvyorganizationservice:3", "idle", 1, "o-123asdfasd", "o-123asdfasd", "Test Org"},
		{4, "Development", "AWS", "123456789102", "2023-06-26 23:09:22", []string{}, false, "DEFAULT", 2, 143, "2023-07-05 15:59:51", "insightcloudsec_sts_assume_role", "divvyorganizationservice:4", "divvyorganizationservice:4", "idle", 1, "o-123asdfasd", "o-123asdfasd", "Test Org"},
	}

	setup()
	mux.HandleFunc("/v2/public/clouds/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/list_clouds.json"))
	})
	resp, err := client.ListClouds()
	assert.NoError(t, err)

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Validating_%s", tc.name), func(t *testing.T) {
			assert.Equal(t, tc.id, resp[i].ID)
			assert.Equal(t, tc.account_id, resp[i].AccountID)
			assert.Equal(t, tc.name, resp[i].Name)
			assert.Equal(t, tc.type_id, resp[i].TypeID)
			assert.Equal(t, tc.create, resp[i].CreationTime)
			assert.Equal(t, tc.host_assess_status, resp[i].HostAssessmentStatus)
			assert.Equal(t, tc.host_assess_enabled, resp[i].HostAssessmentEnabled)
			assert.Equal(t, tc.status, resp[i].Status)
			assert.Equal(t, tc.badges, resp[i].BadgeCount)
			assert.Equal(t, tc.resources, resp[i].ResourceCount)
			assert.Equal(t, tc.refreshed, resp[i].LastRefereshed)
			assert.Equal(t, tc.role, resp[i].RoleArn)
			assert.Equal(t, tc.group_resource_id, resp[i].GroupResourceID)
			assert.Equal(t, tc.resource_id, resp[i].ResourceID)
			assert.Equal(t, tc.edh_role, resp[i].EDHRole)
			assert.Equal(t, tc.strategy_id, resp[i].StrategyID)
			assert.Equal(t, tc.org_id, resp[i].CloudOrganizationID)
			assert.Equal(t, tc.org_id, resp[i].CloudOrganizationDomainName)
			assert.Equal(t, tc.org_nickname, resp[i].CloudOrganizationNickname)
		})
	}
	teardown()
}
