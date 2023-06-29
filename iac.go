package insightcloudsecClient

import (
	"fmt"
	"net/http"
)

func (c *Client) ListIACConfigs() ([]IACConfig, error) {
	// Lists IAC Configurations and basic stats

	var resp []IACConfig
	err := c.makeRequest(http.MethodGet, "/v3/iac/configs", nil, &resp)
	return resp, err
}

func (c *Client) InitiateIACScan(details Scan, simplified bool) (ScanResults, error) {
	// Initates an IAC Scan and returns results as detailed or simplified

	url := "/v3/iac/scan"
	if simplified {
		url = fmt.Sprintf("%s?readable=true", url)
	}

	var resp ScanResults
	err := c.makeRequest(http.MethodPost, url, details, &resp)
	return resp, err
}

func (c *Client) GetIACScan(build_id int, simplified bool) (ScanResults, error) {
	// Returns the scan results for given build_id

	url := fmt.Sprintf("/v3/iac/scans/%d", build_id)
	if simplified {
		url = fmt.Sprintf("%s?readable=true", url)
	}

	var resp ScanResults
	err := c.makeRequest(http.MethodGet, url, nil, &resp)
	return resp, err
}
