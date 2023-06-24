package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroups_ListGroups(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/groups/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("groups/list_groups_response.json"))
	})

	groups, err := client.ListGroups()
	assert.NoError(t, err)
	assert.Equal(t, 20, groups.Groups[0].ID)
	assert.Equal(t, 21, groups.Groups[1].ID)
	assert.Equal(t, "divvyusergroup:20", groups.Groups[0].ResourceID)
	assert.Equal(t, "divvyusergroup:21", groups.Groups[1].ResourceID)
	assert.Equal(t, "My Fun Users", groups.Groups[0].Name)
	assert.Equal(t, "Test UserGroup", groups.Groups[1].Name)
	assert.Equal(t, 10, groups.Groups[0].Users)
	assert.Equal(t, 1, groups.Groups[1].Users)
	assert.Equal(t, 2, groups.Groups[0].Roles)
	assert.Equal(t, 0, groups.Groups[1].Roles)
	assert.True(t, groups.Groups[0].EntitlementsConfigured)
	assert.False(t, groups.Groups[1].EntitlementsConfigured)
	teardown()
}

func TestGroups_GetGroupByID(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/groups/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("groups/list_groups_response.json"))
	})

	// By int
	group1, err := client.GetGroupByID(20)
	assert.NoError(t, err)
	assert.Equal(t, "My Fun Users", group1.Name)
	// By string
	group2, err := client.GetGroupByID("divvyusergroup:21")
	assert.NoError(t, err)
	assert.Equal(t, "Test UserGroup", group2.Name)
	// Invalid string
	_, err = client.GetGroupByID("divvvyusergrouppppp:21")
	assert.Error(t, err)
	// Invalid type
	var a interface{}
	_, err = client.GetGroupByID(a)
	assert.Error(t, err)
	teardown()
}

func TestGroups_CreateGroup(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_DeleteGroup(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_AddGroupUsers(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_UpdateAllGroupUsers(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_DeleteGroupUser(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_ListGroupUsers(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_ListGroupRoles(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_UpdateGroupRoles(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_ListGroupEntitlements(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_SetEntitelments(t *testing.T) {
	setup()

	teardown()
}

func TestGroups_ListUserEntitlement(t *testing.T) {
	setup()

	teardown()
}
