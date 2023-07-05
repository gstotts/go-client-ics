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
	assert.NotEqual(t, len(resp), 0)

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Validating_%s", tc.cloud_type_id), func(t *testing.T) {
			assert.Equal(t, tc.cloud_type_id, resp[i].ID)
			assert.Equal(t, tc.cloud_type_name, resp[i].Name)
			assert.Equal(t, tc.cloud_access, resp[i].Access)
		})
	}
	teardown()
}
