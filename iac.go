package insightcloudsecClient

import "net/http"

func (c *Client) ListIACConfigs() ([]IACConfig, error) {
	// Lists IAC Configurations and basic stats

	var resp []IACConfig
	err := c.makeRequest(http.MethodGet, "/v3/iac/configs", nil, &resp)
	return resp, err
}
