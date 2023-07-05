package insightcloudsecClient

import (
	"net/http"
)

func (c *Client) ListCloudTypes() (CloudTypes, error) {
	// Returns a list of all cloud types

	var cloudTypes CloudTypes
	err := c.makeRequest(http.MethodGet, "/v2/public/cloudtypes/list", nil, cloudTypes)
	return cloudTypes, err
}
