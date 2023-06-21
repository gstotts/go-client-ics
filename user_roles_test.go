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

	teardown()
}

func TestUserRoles_GetRoleByID(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/roles/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'DELETE', got %s", r.Method)
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
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'DELETE', got %s", r.Method)
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
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'DELETE', got %s", r.Method)
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
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'DELETE', got %s", r.Method)
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
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'DELETE', got %s", r.Method)
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
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'DELETE', got %s", r.Method)
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
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/list_roles_response.json"))
	})
	roles, err := client.ListRoles()
	assert.NoError(t, err)
	assert.Len(t, roles.Roles, 3)
	assert.Equal(t, "divvyrole:1:20", roles.Roles[0].ResourceID)
	assert.Equal(t, "divvyrole:1:21", roles.Roles[1].ResourceID)
	assert.Equal(t, "divvyrole:1:22", roles.Roles[2].ResourceID)
	assert.Equal(t, "Random Role 213", roles.Roles[0].Name)
	assert.Equal(t, "Global Read Only", roles.Roles[1].Name)
	assert.Equal(t, "Updated Name 5", roles.Roles[2].Name)
	assert.Equal(t, "Allows stuff", roles.Roles[0].Description)
	assert.Equal(t, "Allows only read for all", roles.Roles[1].Description)
	assert.Equal(t, "Updated Description 5", roles.Roles[2].Description)
	assert.Equal(t, "AND", roles.Roles[0].BadgeFilterOperator)
	assert.Equal(t, "AND", roles.Roles[1].BadgeFilterOperator)
	assert.Equal(t, "OR", roles.Roles[2].BadgeFilterOperator)
	assert.True(t, roles.Roles[0].AllPermissions)
	assert.False(t, roles.Roles[1].AllPermissions)
	assert.False(t, roles.Roles[2].AllPermissions)
	assert.True(t, roles.Roles[0].GlobalScope)
	assert.True(t, roles.Roles[1].GlobalScope)
	assert.False(t, roles.Roles[2].GlobalScope)
	assert.False(t, roles.Roles[0].View)
	assert.True(t, roles.Roles[1].View)
	assert.False(t, roles.Roles[2].View)
	assert.False(t, roles.Roles[0].Provision)
	assert.False(t, roles.Roles[1].Provision)
	assert.True(t, roles.Roles[2].Provision)
	assert.False(t, roles.Roles[0].Manage)
	assert.True(t, roles.Roles[1].Manage)
	assert.False(t, roles.Roles[2].Manage)
	assert.False(t, roles.Roles[0].Delete)
	assert.False(t, roles.Roles[1].Delete)
	assert.True(t, roles.Roles[2].Delete)
	assert.False(t, roles.Roles[0].DeleteCloud)
	assert.False(t, roles.Roles[1].DeleteCloud)
	assert.True(t, roles.Roles[2].DeleteCloud)
	assert.Empty(t, roles.Roles[0].CloudScopes)
	assert.Empty(t, roles.Roles[2].CloudScopes)
	assert.Empty(t, roles.Roles[0].ResourceGroupScopes)
	assert.Empty(t, roles.Roles[2].ResourceGroupScopes)
	assert.Empty(t, roles.Roles[0].BadgeScopes)
	assert.Empty(t, roles.Roles[2].BadgeScopes)
	assert.ElementsMatch(t, []string{"divvyorganizationservice:1"}, roles.Roles[1].CloudScopes)
	assert.ElementsMatch(t, []string{"resourcegroup:1:"}, roles.Roles[1].ResourceGroupScopes)
	assert.ElementsMatch(t, []string{"cloud_type:GCP"}, roles.Roles[1].BadgeScopes)
	assert.ElementsMatch(t, []string{"divvyusergroup:20"}, roles.Roles[1].Groups)
	assert.Empty(t, roles.Roles[2].CloudScopes)
	assert.Empty(t, roles.Roles[2].ResourceGroupScopes)
	assert.Empty(t, roles.Roles[2].BadgeScopes)
	assert.Empty(t, roles.Roles[0].Groups)

	teardown()
}

func TestUserRoles_UpdateRoleScope(t *testing.T) {
	setup()

	teardown()
}

func TestUserRoles_UpdateRoleUserGroups(t *testing.T) {
	setup()

	teardown()
}
