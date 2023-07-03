package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) ListInsights() ([]Insight, error) {

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
