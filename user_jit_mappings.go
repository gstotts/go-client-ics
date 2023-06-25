package insightcloudsecClient

import (
	"fmt"
	"net/http"
)

func (c *Client) ListGroupMappings(auth_server_id int) ([]GroupMapping, error) {
	// Lists the group mapping list associated with the authentication server

	var resp []GroupMapping
	err := c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/prototype/authenticationserver/%d/group_mapping", auth_server_id), nil, &resp)
	return resp, err
}

func (c *Client) AddGroupMapping(auth_server_id int, mapping []GroupMapping) error {
	// Adds a new group mapping to the mapping list

	return c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/authenticationserver/%d/insert_group_mapping", auth_server_id), mapping, nil)
}

func (c *Client) UpdateAllGroupMappings(auth_server_id int, mapping []GroupMapping) error {
	// Overwrites the group mapping with the given mapping list

	return c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/prototype/authenticationserver/%d/update_group_mapping", auth_server_id), mapping, nil)
}

func (c *Client) DeleteGroupMapping(auth_server_id int, mapping []GroupMapping) error {
	// Deletes an existing group mapping.  Non-existent mappings are ignored.

	return c.makeRequest(http.MethodDelete, fmt.Sprintf("/v2/prototype/authenticationserver/%d/delete_group_mapping", auth_server_id), mapping, nil)
}
