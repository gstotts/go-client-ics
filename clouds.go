package insightcloudsecClient

import (
	"fmt"
	"net/http"
)

func (c *Client) ListCloudTypes() ([]CloudType, error) {
	// Returns a list of all cloud types

	var cloudTypes CloudTypes
	err := c.makeRequest(http.MethodGet, "/v2/public/cloudtypes/list", nil, cloudTypes)
	fmt.Println(cloudTypes)
	return cloudTypes.CloudTypes, err
}
