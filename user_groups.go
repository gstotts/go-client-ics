package insightcloudsecClient

import "net/http"

func (c *Client) ListGroups() (Groups, error) {
	// Returns a list of user groups

	var resp Groups
	err := c.makeRequest(http.MethodGet, "/v2/prototype/groups/list", nil, &resp)
	return resp, err
}

func (c *Client) DeleteGroup(group_resource_id string) error {
	return nil
}
