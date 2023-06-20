package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticationServers_ListAuthenticationServers(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/authenticationservers/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("authentication_servers/auth_servers_list_response.json"))
	})
	resp, err := client.ListAuthetnicationServers()
	assert.NoError(t, err)
	assert.Equal(t, 1, resp.Servers[0].ID)
	assert.Equal(t, "saml_server", resp.Servers[0].Name)
	assert.Equal(t, "mysaml.mydomain.org", resp.Servers[0].Host)
	assert.Equal(t, 123, resp.Servers[0].Port)
	assert.Equal(t, 1, resp.Servers[0].Secure)
	assert.Equal(t, "saml", resp.Servers[0].Type)
	assert.Equal(t, true, resp.Servers[0].GlobalScope)
	assert.Equal(t, 5, resp.Servers[0].MappedGroups)
	teardown()
}
