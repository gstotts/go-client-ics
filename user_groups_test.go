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

	testCases := []struct {
		id           int
		resource_id  string
		name         string
		users        int
		roles        int
		entitlements bool
	}{
		{20, "divvyusergroup:20", "My Fun Users", 10, 2, true},
		{21, "divvyusergroup:21", "Test UserGroup", 1, 0, false},
	}

	groups, err := client.ListGroups()
	assert.NoError(t, err)

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test Group %d", i), func(t *testing.T) {
			assert.Equal(t, tc.id, groups.Groups[i].ID)
			assert.Equal(t, tc.resource_id, groups.Groups[i].ResourceID)
			assert.Equal(t, tc.name, groups.Groups[i].Name)
			assert.Equal(t, tc.users, groups.Groups[i].Users)
			assert.Equal(t, tc.roles, groups.Groups[i].Roles)
			assert.Equal(t, tc.entitlements, groups.Groups[i].EntitlementsConfigured)
		})
	}
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

	testCases := []struct {
		test_name string
		id        any
		name      string
		err       bool
	}{
		{"Valid Group by Int", 20, "My Fun Users", false},
		{"Valid Group by String", "divvyusergroup:21", "Test UserGroup", false},
		{"Invalid Group by Int", 22, "", true},
		{"Invalid Group by Type", []int{1, 2}, "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			group, err := client.GetGroupByID(tc.id)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.name, group.Name)
		})
	}
	teardown()
}

func TestGroups_CreateGroup(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/group/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("groups/create_group_response.json"))
	})

	group, err := client.CreateGroup("Test UserGroup")
	assert.NoError(t, err)
	assert.Equal(t, "Test UserGroup", group.Name)
	teardown()
}

func TestGroups_DeleteGroup(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/group/divvyusergroup:25/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	testCases := []struct {
		id  string
		err bool
	}{
		{"divvyusergroup:25", false},
		{"divvyusergroup:22", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Delete Test Group: %s", tc.id), func(t *testing.T) {
			if tc.err {
				assert.Error(t, client.DeleteGroup(tc.id))
			} else {
				assert.NoError(t, client.DeleteGroup(tc.id))
			}
		})
	}
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
