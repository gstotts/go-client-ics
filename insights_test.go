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
	testCases := []struct {
		test_name     string
		insight_count int
		err_expected  bool
	}{
		{"Valid_Request", 5, false},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			setup()
			mux.HandleFunc("/v2/public/insights/list", func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, getJSONFile(fmt.Sprintf("insights/list_insights_partial.json")))
			})
			resp, err := client.ListInsights()
			if tc.err_expected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.insight_count, len(resp))
			}
			teardown()
		})

	}
}

func TestInsights_QueryInsights(t *testing.T) {
	testCases := []struct {
		test_name              string
		detail                 bool
		labels                 string
		pack_ids               string
		resource_types         string
		expected_in_results    string
		expected_insight_count int
		err_expected           bool
	}{
		{"Query_Detailed_Security_Instance", true, "security", "", "instance", "Instance Exposing SSH to the Public", 21, false},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			setup()
			mux.HandleFunc("/v2/public/insights/list", func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, fmt.Sprintf("%t", tc.detail), r.URL.Query().Get("detail"))
				if tc.labels != "" {
					assert.Equal(t, tc.labels, r.URL.Query().Get("labels"))
				}
				if tc.pack_ids != "" {
					assert.Equal(t, tc.pack_ids, r.URL.Query().Get("pack_ids"))
				}
				if tc.resource_types != "" {
					assert.Equal(t, tc.resource_types, r.URL.Query().Get("resource_types"))
				}
				assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, getJSONFile(fmt.Sprintf("insights/query_%t_%s_%s_%s.json", tc.detail, tc.labels, tc.pack_ids, tc.resource_types)))
			})
			resp, err := client.QueryInsights(tc.detail, tc.labels, tc.pack_ids, tc.resource_types)
			if tc.err_expected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected_insight_count, len(resp))
				found_expected_insight := false
				for _, data := range resp {
					assert.Contains(t, data.ResourceTypes, tc.resource_types)
					if data.Name == tc.expected_in_results {
						found_expected_insight = true
					}
				}
				assert.True(t, found_expected_insight)
			}
			teardown()
		})

	}
}

func TestInsights_ListFilters(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/insights/filter-registry", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("insights/list_filters_partial.json"))
	})
	filters, err := client.ListFilters()
	assert.NoError(t, err)
	for name, filter := range filters {
		assert.Equal(t, "divvy.query.access_analyzer_finding_count", name)
		assert.Equal(t, "divvy.query.access_analyzer_finding_count", filter.ID)
		assert.Equal(t, "Access Analyzer Finding Count By Type", filter.Name)
		assert.Equal(t, []string{"accessanalyzer"}, filter.SupportedResources)
		assert.Equal(t, []string{"AWS", "AWS_GOV", "AWS_CHINA"}, filter.SupportedClouds)
		assert.False(t, filter.SupportsCommon)
		assert.Equal(t, 3, len(filter.SettingsConfig))
	}
	teardown()
}
