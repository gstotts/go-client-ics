package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) GetInsight(insight_id int, insight_source string) (Insight, error) {
	// Returns the insight of given id and source

	if strings.ToLower(insight_source) != "custom" && strings.ToLower(insight_source) != "backoffice" {
		return Insight{}, fmt.Errorf("insight_source must be one of custom or backoffice, got: %s", insight_source)
	}
	var resp Insight
	err := c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/insights/%d/%s", insight_id, strings.ToLower(insight_source)), nil, &resp)
	return resp, err
}

func (c *Client) ListInsights() ([]Insight, error) {
	//Returns a list of all insights

	var resp []Insight
	err := c.makeRequest(http.MethodGet, "/v2/public/insights/list", nil, &resp)
	return resp, err
}

func (c *Client) QueryInsights(detail bool, labels, pack_ids, resource_types string) ([]Insight, error) {
	// Returns a list of insights based on given parameters.  Empty strings will omit a specific parameter.

	path := fmt.Sprintf("/v2/public/insights/list?detail=%t", detail)
	if labels != "" {
		path = fmt.Sprintf("%s&labels=%s", path, url.PathEscape(labels))
	}
	if pack_ids != "" {
		path = fmt.Sprintf("%s&pack_ids=%s", path, url.PathEscape(pack_ids))
	}
	if resource_types != "" {
		path = fmt.Sprintf("%s&resource_types=%s", path, resource_types)
	}
	var resp []Insight
	err := c.makeRequest(http.MethodGet, path, nil, &resp)
	return resp, err
}

func (c *Client) ListFilters() (map[string]InsightFilter, error) {
	// Lists all filters available for use in Insights and the config details

	var filterRegistry map[string]InsightFilter
	err := c.makeRequest(http.MethodGet, "/v2/public/insights/filter-registry", nil, &filterRegistry)
	return filterRegistry, err

}
