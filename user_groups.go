package insightcloudsecClient

import "net/http"

func (g *Groups) List() (Groups, error) {
	// Returns a list of user groups

	var resp Groups
	err := g.client.makeRequest(http.MethodGet, "/v2/prototype/groups/list", nil, &resp)
	return resp, err
}

func (c *Client) DeleteGroup(group_resource_id string) error {
	return nil
}
