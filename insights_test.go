package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsights_GetInsight(t *testing.T) {
	testCases := []struct {
		test_name    string
		insight_id   int
		source       string
		err_expected bool
	}{
		{"Valid_Insight_Request", 14, "backoffice", false},
		{"Invalid_Insight_Source", 14, "office", true},
		{"Invalid_Insight_ID", 44, "backoffice", true},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			setup()
			mux.HandleFunc(fmt.Sprintf("/v2/public/insights/%d/%s", tc.insight_id, tc.source), func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, getJSONFile(fmt.Sprintf("insights/get_insight_%s_%d.json", tc.source, tc.insight_id)))
			})
			resp, err := client.GetInsight(tc.insight_id, tc.source)
			if tc.err_expected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.insight_id, resp.ID)
			}
			teardown()
		})

	}
}

func TestInsights_ListInsights(t *testing.T) {

}

func TestInsights_QueryInsights(t *testing.T) {}
