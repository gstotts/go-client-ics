package insightcloudsecClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRoles_CreateRole(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/role/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))
	})

	resp, err := client.CreateRole(Role{
		Name:           "Test Role",
		Description:    "Role defines permissions to scopes",
		AllPermissions: false,
		View:           false,
		Provision:      true,
		Manage:         false,
		Delete:         true,
		AddCloud:       false,
		DeleteCloud:    true,
		GlobalScope:    false,
	})
	assert.NoError(t, err)
	assert.Equal(t, "divvyrole:1:24", resp.ResourceID)
	assert.Equal(t, "Test Role", resp.Name)
	assert.Equal(t, "Role defines permissions to scopes", resp.Description)
	assert.False(t, resp.AllPermissions)
	assert.False(t, resp.View)
	assert.True(t, resp.Provision)
	assert.False(t, resp.Manage)
	assert.True(t, resp.Delete)
	assert.False(t, resp.AddCloud)
	assert.True(t, resp.DeleteCloud)
	assert.False(t, resp.GlobalScope)
	teardown()
}

func TestUserRoles_GetRoleByID(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/roles/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/list_roles_response.json"))
	})

	resp, err := client.GetRoleByID("divvyrole:1:21")
	assert.NoError(t, err)
	assert.Equal(t, "divvyrole:1:21", resp.ResourceID)
	assert.Equal(t, "Global Read Only", resp.Name)
	teardown()
}

func TestUserRoles_UpdateRole(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/role/divvyrole:1:24/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.NotNil(t, r.Body)
		req_body := Role{}
		err := json.NewDecoder(r.Body).Decode(&req_body)
		assert.NoError(t, err)
		assert.Empty(t, req_body.ResourceID)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))

	})
	resp, err := client.UpdateRole("divvyrole:1:24", Role{
		Name:           "Test Role",
		Description:    "Role defines permissions to scopes",
		AllPermissions: false,
		View:           false,
		Provision:      true,
		Manage:         false,
		Delete:         true,
		AddCloud:       false,
		DeleteCloud:    true,
		GlobalScope:    false,
		ResourceID:     "divvyrole:1:24",
	})
	assert.NoError(t, err)
	assert.Equal(t, "divvyrole:1:24", resp.ResourceID)
	assert.Equal(t, "Test Role", resp.Name)
	assert.Equal(t, "Role defines permissions to scopes", resp.Description)
	assert.False(t, resp.AllPermissions)
	assert.False(t, resp.View)
	assert.True(t, resp.Provision)
	assert.False(t, resp.Manage)
	assert.True(t, resp.Delete)
	assert.False(t, resp.AddCloud)
	assert.True(t, resp.DeleteCloud)
	assert.False(t, resp.GlobalScope)
	teardown()
}

func TestUserRoles_UpdateRole_InvalidRequestBody_BadgeFilterOperator(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/role/divvyrole:1:9/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.NotNil(t, r.Body)
		req_body := Role{}
		err := json.NewDecoder(r.Body).Decode(&req_body)
		assert.NoError(t, err)
		assert.Empty(t, req_body.ResourceID)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))

	})
	_, err := client.UpdateRole("divvyrole:1:9", Role{
		Name:                "Test Role",
		Description:         "Role defines permissions to scopes",
		AllPermissions:      true,
		View:                false,
		Provision:           false,
		Manage:              false,
		Delete:              false,
		AddCloud:            false,
		DeleteCloud:         false,
		GlobalScope:         false,
		BadgeFilterOperator: "AND",
	})
	assert.Error(t, err)
	teardown()
}

func TestUserRoles_UpdateRole_InvalidRequestBody_Groups(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/role/divvyrole:1:9/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.NotNil(t, r.Body)
		req_body := Role{}
		err := json.NewDecoder(r.Body).Decode(&req_body)
		assert.NoError(t, err)
		assert.Empty(t, req_body.ResourceID)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))

	})
	_, err := client.UpdateRole("divvyrole:1:9", Role{
		Name:           "Test Role",
		Description:    "Role defines permissions to scopes",
		AllPermissions: true,
		View:           false,
		Provision:      false,
		Manage:         false,
		Delete:         false,
		AddCloud:       false,
		DeleteCloud:    false,
		GlobalScope:    false,
		ResourceID:     "divvyrole:1:9",
		Groups:         []string{"divvyusergroup:10"},
	})
	assert.Error(t, err)
	teardown()
}

func TestUserRoles_UpdateRole_InvalidRequestBody_ResourceGroupScopes(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/role/divvyrole:1:9/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.NotNil(t, r.Body)
		req_body := Role{}
		err := json.NewDecoder(r.Body).Decode(&req_body)
		assert.NoError(t, err)
		assert.Empty(t, req_body.ResourceID)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))

	})
	_, err := client.UpdateRole("divvyrole:1:9", Role{
		Name:                "Test Role",
		Description:         "Role defines permissions to scopes",
		AllPermissions:      true,
		View:                false,
		Provision:           false,
		Manage:              false,
		Delete:              false,
		AddCloud:            false,
		DeleteCloud:         false,
		GlobalScope:         false,
		ResourceGroupScopes: []string{"resourcegroup:1:"},
	})
	assert.Error(t, err)
	teardown()
}

func TestUserRoles_UpdateRole_InvalidRequestBody_BadgeScopes(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/role/divvyrole:1:9/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.NotNil(t, r.Body)
		req_body := Role{}
		err := json.NewDecoder(r.Body).Decode(&req_body)
		assert.NoError(t, err)
		assert.Empty(t, req_body.ResourceID)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))

	})
	_, err := client.UpdateRole("divvyrole:1:9", Role{
		Name:           "Test Role",
		Description:    "Role defines permissions to scopes",
		AllPermissions: true,
		View:           false,
		Provision:      false,
		Manage:         false,
		Delete:         false,
		AddCloud:       false,
		DeleteCloud:    false,
		GlobalScope:    false,
		BadgeScopes:    []string{"cloud_type:GCP"},
	})
	assert.Error(t, err)
	teardown()
}

func TestUserRoles_ListRoles(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/roles/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/list_roles_response.json"))
	})

	testCases := []struct {
		resource_id           string
		name                  string
		desc                  string
		badgeoperator         string
		allperms              bool
		view                  bool
		provision             bool
		manage                bool
		delete                bool
		add_cloud             bool
		delete_cloud          bool
		global_scope          bool
		cloud_scopes          []string
		resource_group_scopes []string
		badge_Scopes          []string
		groups                []string
	}{
		{"divvyrole:1:20", "Random Role 213", "Allows stuff", "AND", true, false, false, false, false, true, false, true, []string{}, []string{}, []string{}, []string{}},
		{"divvyrole:1:21", "Global Read Only", "Allows only read for all", "AND", false, true, false, true, false, false, false, true, []string{"divvyorganizationservice:1"}, []string{"resourcegroup:1:"}, []string{"cloud_type:GCP"}, []string{"divvyusergroup:20"}},
		{"divvyrole:1:22", "Updated Name 5", "Updated Description 5", "OR", false, false, true, false, true, true, true, false, []string{}, []string{}, []string{}, []string{"divvyusergroup:20"}},
	}

	roles, err := client.ListRoles()
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 3)

	for i, tc := range testCases {
		t.Run(tc.resource_id, func(t *testing.T) {
			assert.Equal(t, tc.resource_id, roles.Roles[i].ResourceID)
			assert.Equal(t, tc.name, roles.Roles[i].Name)
			assert.Equal(t, tc.desc, roles.Roles[i].Description)
			assert.Equal(t, tc.badgeoperator, roles.Roles[i].BadgeFilterOperator)
			assert.Equal(t, tc.allperms, roles.Roles[i].AllPermissions)
			assert.Equal(t, tc.view, roles.Roles[i].View)
			assert.Equal(t, tc.provision, roles.Roles[i].Provision)
			assert.Equal(t, tc.manage, roles.Roles[i].Manage)
			assert.Equal(t, tc.delete, roles.Roles[i].Delete)
			assert.Equal(t, tc.add_cloud, roles.Roles[i].AddCloud)
			assert.Equal(t, tc.delete_cloud, roles.Roles[i].DeleteCloud)
			assert.Equal(t, tc.global_scope, roles.Roles[i].GlobalScope)
			assert.Equal(t, tc.cloud_scopes, roles.Roles[i].CloudScopes)
			assert.Equal(t, tc.resource_group_scopes, roles.Roles[i].ResourceGroupScopes)
			assert.Equal(t, tc.badge_Scopes, roles.Roles[i].BadgeScopes)
			assert.Equal(t, tc.groups, roles.Roles[i].Groups)

		})
	}
	teardown()
}

func TestUserRoles_UpdateRoleScope(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/roles/divvyrole:1:24/scope/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.NotNil(t, r.Body)
		var req_body rolesUpdateScopeRequest
		err := json.NewDecoder(r.Body).Decode(&req_body)
		assert.NoError(t, err)
		assert.Equal(t, []string{"resourcegroup:1", "divvyorganizationservice:1"}, req_body.ResourceIDs)
		assert.Equal(t, []string{"resourcegroup:2"}, req_body.DeprecatedResourceIDs)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))
	})
	resp, err := client.UpdateRoleScope("divvyrole:1:24", []string{"resourcegroup:1", "divvyorganizationservice:1"}, []string{"resourcegroup:2"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"resourcegroup:1:"}, resp.ResourceGroupScopes)
	assert.Equal(t, []string{"divvyorganizationservice:1"}, resp.CloudScopes)
	teardown()
}

func TestUserRoles_UpdateRoleUserGroups(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/roles/divvyrole:1:24/groups/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.NotNil(t, r.Body)
		var req_body rolesUpdateUserGroupsRequest
		err := json.NewDecoder(r.Body).Decode(&req_body)
		assert.NoError(t, err)
		assert.Equal(t, []string{"divvyusergroup:20"}, req_body.GroupResourceIDs)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))
	})
	resp, err := client.UpdateRoleUserGroups("divvyrole:1:24", []string{"divvyusergroup:20"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"divvyusergroup:20"}, resp.Groups)
	teardown()
}

func TestUserRoles_UpdateRoleBadges(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/roles/divvyrole:1:24/AND/badges/update", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.NotNil(t, r.Body)
		var req_body Badges
		err := json.NewDecoder(r.Body).Decode(&req_body)
		assert.NoError(t, err)
		assert.Equal(t, Badge{Key: "cloud_type", Value: "GCP"}, req_body.Badges[0])
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))
	})
	resp, err := client.UpdateRoleBadges("divvyrole:1:24", "AND", []Badge{{Key: "cloud_type", Value: "GCP"}})
	assert.NoError(t, err)
	assert.Equal(t, []string{"cloud_type:GCP"}, resp.BadgeScopes)
	teardown()
}
