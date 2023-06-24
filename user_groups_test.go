package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroups_ListGroups(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/group/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("roles/role_response.json"))
	})

	teardown()
}

func TestGroups_GetGroupByID(t *testing.T) {
	setup()

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
