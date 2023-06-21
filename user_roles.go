package insightcloudsecClient

import (
	"fmt"
	"net/http"
)

func (c *Client) CreateRole(r Role) (Role, error) {
	// Creates a role with the given details from the Role struct type.
	// Badges, Resources and Groups cannot be attached until after creation.

	if r.ResourceID != "" {
		return Role{}, fmt.Errorf("resource id must not be set.  current value: %s", r.ResourceID)
	}

	resp := Role{}
	err := c.makeRequest(http.MethodPost, "/v2/public/role/create", r, &resp)

	return resp, err
}

func (c *Client) GetRoleByID(resource_id string) (Role, error) {
	// Returns a specific role of the given role_id

	roles, err := c.ListRoles()
	if err != nil {
		return Role{}, fmt.Errorf("error retrieving list of roles:\n%s", err)
	}

	for _, r := range roles.Roles {
		if r.ResourceID == resource_id {
			return r, nil
		}
	}

	return Role{}, fmt.Errorf("unable to find role with resource_id: %s", resource_id)
}

func (c *Client) UpdateRole(role_resource_id string, r Role) (Role, error) {
	// Updates the existing role at the given resource_id with the config of r Role type struct

	resp := Role{}
	err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/role/%s/update", role_resource_id), r, &resp)
	return resp, err
}

func (c *Client) ListRoles() (Roles, error) {
	// Returns a list of user roles from the InsightCloudSec API

	resp := Roles{}
	err := c.makeRequest(http.MethodPost, "/v2/public/roles/list", nil, &resp)
	return resp, err
}

func (c *Client) UpdateRoleScope(role_resource_id string, resource_ids, deprecated_resource_ids []string) (Role, error) {
	// Allows you to update a role's scope for resource_ids

	// Build payload
	data := map[string]interface{}{
		"resource_ids":            resource_ids,
		"deprecated_resource_ids": deprecated_resource_ids,
	}

	resp := Role{}
	err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/roles/%s/scope/update", role_resource_id), data, &resp)
	return resp, err
}

func (c *Client) UpdateRoleUserGroups(role_resource_id string, group_ids []string) (Role, error) {
	// Allows you to update a role's scope for user groups

	// Build payload
	data := map[string]interface{}{
		"group_resource_ids": group_ids,
	}

	resp := Role{}
	err := c.makeRequest(http.MethodGet, fmt.Sprintf("/v2/public/roles/%s/groups/update", role_resource_id), data, &resp)
	return resp, err

}
