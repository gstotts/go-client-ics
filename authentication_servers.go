package insightcloudsecClient

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListAuthetnicationServers() (AuthenticationServers, error) {
	// Returns a list of authentication servers from the InsightCloudSec API

	// Make Request
	body, err := c.makeRequest(http.MethodPost, "/v2/prototype/authenticationservers/list", nil)
	if err != nil {
		return AuthenticationServers{}, err
	}

	// Unmarshal Response
	resp := AuthenticationServers{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return AuthenticationServers{}, err
	}

	return resp, nil
}
