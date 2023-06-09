package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth_Login_NoUser(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/login", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("auth/auth_login_response.json"))
	})

	client.Auth = AuthStruct{
		Username: "",
		Password: "1234#@!",
	}

	client.APIKey = ""

	_, err := client.Login()
	assert.Error(t, err)
	teardown()
}

func TestAuth_Login_NoPassword(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/login", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("auth/auth_login_response.json"))
	})

	client.Auth = AuthStruct{
		Username: "han_solo",
		Password: "",
	}

	client.APIKey = ""

	_, err := client.Login()
	assert.Error(t, err)
	teardown()
}

func TestAuth_Login(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/login", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("auth/auth_login_response.json"))
	})

	client.Auth = AuthStruct{
		Username: "han_solo",
		Password: "1234#@!",
	}

	client.APIKey = ""

	resp, err := client.Login()
	assert.NoError(t, err)
	assert.Equal(t, 123, resp.UserID)
	assert.Equal(t, "han_solo", resp.Name)
	assert.Equal(t, "han_solo@mfalcon.com", resp.Email)
	assert.Equal(t, "1234:ab:cd:ef:gh:12:24:a1A:bb:3c", resp.CustomerID)
	assert.Equal(t, "1fasdf1324asdfasdfasdf11246", resp.SessionID)
	assert.Equal(t, 720, resp.Timeout)
	assert.Equal(t, false, resp.AuthPluginExists)
	assert.Equal(t, true, resp.DomainAdmin)
	assert.Equal(t, false, resp.DomainViewer)
	teardown()
}

func TestAuth_CreateAPIKey(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/apikey/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("auth/auth_apikey_response.json"))
	})

	resp, err := client.CreateAPIKey(44)
	assert.NoError(t, err)
	assert.Equal(t, "123456789AbCdEfGhIjKlMnOpQrStUvWxYz987654321", resp)
	teardown()
}
