package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"time"
)

func (c *Client) ListPrincipalActivity(principal_resource_id, start, end string) (PrincipalActivity, error) {
	// Lists the activity (policy / permission actions) for a given principal

	if start == "" || end == "" {
		return PrincipalActivity{}, fmt.Errorf("start and end must be provided (formatted string as YYYY-MM-DD)")
	}

	// Validate start is proper format
	if !validateDateFormats(start, end) {
		return PrincipalActivity{}, fmt.Errorf("start/end time is not of format YYYY-MM-DD.  Got %s and %s", start, end)
	}

	path_and_query := fmt.Sprintf("/v3/lpa/principals/%s/actions?start=%s&end=%s", principal_resource_id, start, end)

	// Make Request
	var resp PrincipalActivity
	err := c.makeRequest(http.MethodGet, path_and_query, nil, &resp)

	return resp, err
}

func (c *Client) ListPrincipalPermissions(principal_resource_id string) (PrincipalPermissions, error) {
}

func validateDateFormats(start, end string) bool {
	_, err := time.Parse("2006-01-02", start)
	_, err2 := time.Parse("2006-01-02", end)

	if err != nil || err2 != nil {
		return true
	}

	return false
}
