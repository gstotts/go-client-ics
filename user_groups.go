package insightcloudsecClient

import (
	"fmt"
	"net/http"
)

func (c *Client) ListGroups() (Groups, error) {
	// Returns a list of user groups

	var resp Groups
	err := c.makeRequest(http.MethodGet, "/v2/prototype/groups/list", nil, &resp)
	return resp, err
}

func (c *Client) GetGroupByID(group_id any) (Group, error) {
	// Returns a group of given group_id (either int of id or string of resource_id)

	groups, err := c.ListGroups()
	if err != nil {
		return Group{}, fmt.Errorf("unable to retrive groups: %s", err)
	}

	var result Group
	switch g := group_id.(type) {
	case int:
		result, err = checkGroupsByID(g, groups)
	case string:
		result, err = checkGroupsByResourceID(g, groups)
	default:
		err = fmt.Errorf("given id must be int of group_id or string of resource_id, got %T", g)
	}

	return result, err
}

func checkGroupsByID(id int, groups Groups) (Group, error) {
	for _, group := range groups.Groups {
		if group.ID == id {
			return group, nil
		}
	}

	return Group{}, fmt.Errorf("unable to find group of id: %d", id)
}

func checkGroupsByResourceID(id string, groups Groups) (Group, error) {
	for _, group := range groups.Groups {
		if group.ResourceID == id {
			return group, nil
		}
	}

	return Group{}, fmt.Errorf("unable to find group of resource_id: %s", id)
}

func (c *Client) CreateGroup(group_name string) (Group, error) {
	// Creates a group of given name and returns it

	var resp groupResponse
	err := c.makeRequest(http.MethodPost, "/v2/prototype/group/create", createGroupRequest{Name: group_name}, &resp)
	return resp.Group, err
}

func (c *Client) DeleteGroup(group_resource_id string) error {
	// Deletes the group of given resource_id

	return c.makeRequest(http.MethodDelete, fmt.Sprintf("/v2/prototype/group/%s/delete", group_resource_id), nil, nil)
}

func (c *Client) AddGroupUsers(group_resource_id string, users []string) (Group, error) {
	// Adds the slice of given users to the group

	var resp groupResponse
	err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/group/%s/users/add", group_resource_id), addUsersToGroupRequest{UserResourceIDs: users}, &resp)
	return resp.Group, err
}

func (c *Client) UpdateAllGroupUsers(group_resource_id string, users []string) (Group, error) {
	// Replaces current users in the group with the given list.  You must be an Org Admin to utilize.

	current_user, err := c.CurrentUserInfo()
	if err != nil {
		return Group{}, fmt.Errorf("error validating user is org admin prior:\n%s", err)
	}

	if !(current_user.OrganizationAdmin) {
		return Group{}, fmt.Errorf("user must be org admin to update all group members")
	}

	var resp groupResponse
	err = c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/group/%s/users/update", group_resource_id), addUsersToGroupRequest{UserResourceIDs: users}, &resp)
	return resp.Group, err
}

func (c *Client) DeleteGroupUser(group_resource_id, user_resource_id string) (Group, error) {
	// Deletes user of user_resource_id from group of group_resource_id

	var resp groupResponse
	err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/group/%s/user/remove", group_resource_id), deleteUserFromGroupRequest{UserResourceID: user_resource_id}, &resp)
	return resp.Group, err
}

func (c *Client) ListGroupUsers(group_resource_id string) (Users, error) {
	// Returns all users of the given group

	var resp Users
	err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/group/%s/users/list", group_resource_id), nil, &resp)
	return resp, err
}

func (c *Client) ListGroupRoles(group_resource_id string) (Roles, error) {
	// Returns the roles associated with the group

	var resp Roles
	err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/group/%s/roles/list", group_resource_id), nil, &resp)
	return resp, err
}

// func (c *Client) UpdateGroupRoles(group_resource_id string, role_resource_ids []string) (Roles, error) {}

// func (c *Client) ListGroupEntitlements(group_resource_id string) (Entitlements, error) {}

// func (c *Client) SetEntitlements() (Entitlements, error) {}

// func (c *Client) ListUserEntitlement(user_resource_id, module_name string) (Entitlement, error) {}
