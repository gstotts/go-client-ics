package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJIT_ListGroupMappings(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/authenticationserver/1/group_mapping", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("jit_group_mappings/generic_group_mapping.json"))
	})
	mapping, err := client.ListGroupMappings(1)
	assert.NoError(t, err)
	assert.Equal(t, "Domain Admins", mapping[0].LocalName)
	assert.Equal(t, "Domain Viewers", mapping[1].LocalName)
	assert.Equal(t, "Organization Admins", mapping[2].LocalName)
	assert.Equal(t, "Administrators", mapping[0].RemoteName)
	assert.Equal(t, "ReadOnlyAdmins", mapping[1].RemoteName)
	assert.Equal(t, "YourMommasAdmins", mapping[2].RemoteName)
	teardown()
}

func TestJIT_AddGroupMapping(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/authenticationserver/1/insert_group_mapping", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	assert.NoError(t, client.AddGroupMapping(1, []GroupMapping{{
		LocalName:  "MyGroup",
		RemoteName: "StudlyGroup",
	}}))
	teardown()
}

func TestJIT_UpdateGroupMapping(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/authenticationserver/1/update_group_mapping", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	assert.NoError(t, client.UpdateAllGroupMappings(1, []GroupMapping{{
		LocalName:  "MyGroup",
		RemoteName: "StudlyGroup",
	}}))
	teardown()
}

func TestJIT_DeleteGroupMapping(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/authenticationserver/1/delete_group_mapping", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	assert.NoError(t, client.DeleteGroupMapping(1, []GroupMapping{{
		LocalName:  "MyGroup",
		RemoteName: "StudlyGroup",
	}}))
	teardown()
}
