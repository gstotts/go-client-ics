package insightcloudsecClient

import (
	"net/http"
)

func (c *Client) ListCloudTypes() (CloudTypes, error) {
	// Returns a list of all cloud types

	var cloudTypes CloudTypes
	err := c.makeRequest(http.MethodGet, "/v2/public/cloudtypes/list", nil, &cloudTypes)
	return cloudTypes, err
}

func (c *Client) ListClouds() ([]Cloud, error) {
	// List availble configured clouds

	var list Clouds
	err := c.makeRequest(http.MethodGet, "/v2/public/clouds/list", nil, &list)
	return list.Clouds, err
}
