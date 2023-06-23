package insightcloudsecClient

import (
	"net/http"
)

func (c *Client) ListAuthetnicationServers() (AuthenticationServers, error) {
	// Returns a list of authentication servers from the InsightCloudSec API

	resp := AuthenticationServers{}
	err := c.makeRequest(http.MethodPost, "/v2/prototype/authenticationservers/list", nil, &resp)
	return resp, err
}
