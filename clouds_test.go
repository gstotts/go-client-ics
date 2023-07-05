package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClouds_ListCloudTypes(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/cloudtypes/list", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("clouds/list_cloudtypes.json"))
	})
	resp, err := client.ListCloudTypes()
	teardown()
}
