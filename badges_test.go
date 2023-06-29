package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadges_CreateBadge(t *testing.T) {}

func TestBadges_UpdateCloudBadges(t *testing.T) {}

func TestBadges_DeleteBadges(t *testing.T) {}

func TestBadges_ListResourceBadges(t *testing.T) {}

func TestBadges_ListCloudsWithBadges(t *testing.T) {}

func TestBadges_ListResourcesBadgeCount(t *testing.T) {
	testCases := []struct {
		test_name       string
		resource_ids    []string
		expected_counts []int
		test_file       string
		err_expected    bool
	}{
		{"Valid_Request", []string{"divvyorganizationservice:1"}, []int{2}, "badges/list_resources_badge_count.json", false},
		{"ResourceID_Does_Not_Exist", []string{"divvyorganizationservice:12"}, []int{}, "badges/list_resources_badge_count_invalid_id.json", false},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			setup()
			mux.HandleFunc("/v2/public/badges/count", func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, getJSONFile(tc.test_file))
			})

			resp, err := client.ListResourcesBadgeCount(tc.resource_ids)
			if tc.err_expected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				if len(resp) != 0 {
					for z, id := range tc.resource_ids {
						assert.Equal(t, id, resp[z].ResourceID)
						if len(tc.expected_counts) != 0 {
							assert.Equal(t, tc.expected_counts[z], resp[z].Count)
						}
					}
				}
			}
			teardown()
		})
	}
}
