package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUsers_CurrentUserInfo(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/info", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/user_info_response.json"))
	})

	resp, err := client.CurrentUserInfo()
	assert.NoError(t, err)
	assert.Equal(t, "divvyuser:123:", resp.ResourceID)
	assert.Equal(t, "Han Solo", resp.Name)
	assert.Equal(t, 123, resp.UserID)
	assert.Equal(t, false, resp.OrganizationAdmin)
	assert.Equal(t, false, resp.DomainAdmin)
	assert.Equal(t, false, resp.DomainViewer)
	assert.Equal(t, "han_solo@mfalcon.com", resp.EmailAddress)
	assert.Equal(t, "han_solo", resp.Username)
	assert.Equal(t, "Default Organization", resp.OrganizationName)
	assert.Equal(t, 1, resp.OrganizationID)
	assert.Equal(t, true, resp.TwoFactorEnabled)
	assert.Equal(t, false, resp.TwoFactorRequired)
	assert.Equal(t, "local", resp.Auth)
	assert.Equal(t, 3600, resp.SessionTTL)
	assert.Equal(t, "2023/01/01, 01:01:01 UTC", resp.SessionExpiration)
	assert.Equal(t, 3600, resp.SessionTimeoutSeconds)
	assert.Equal(t, 1, resp.AuthenticationServerID)
	assert.Equal(t, false, resp.AuthPluginExists)
	assert.Equal(t, make([]string, 0), resp.NavigationBlacklist)
	assert.Equal(t, "divvy-light-theme", resp.Theme)
	assert.Equal(t, false, resp.RequirePWReset)
	assert.Equal(t, "2022-01-02 12:34:56", resp.CreateDate)
	assert.Equal(t, "1a2bcd3e-45fg-67hi-jklm-8o9p1q2r3st4u", resp.AWSDefaultExternalID)
	assert.Equal(t, false, resp.ApiKeyGenerationAllowed)
	assert.Equal(t, "", resp.Rapid7OrgID)
	teardown()
}

func TestUsers_ListBasicUsers(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/user_list_response.json"))
	})

	resp, err := client.ListBasicUsers()
	assert.NoError(t, err)
	assert.Equal(t, false, resp.Users[0].DomainViewer)
	assert.Equal(t, 1, resp.Users[0].OrganizationID)
	assert.Equal(t, false, resp.Users[0].TwoFactorEnabled)
	assert.Equal(t, false, resp.Users[0].TwoFactorRequired)
	assert.Equal(t, 1, resp.Users[0].Groups)
	assert.Equal(t, make([]string, 0), resp.Users[0].NavigationBlacklist)
	assert.Equal(t, false, resp.Users[0].RequirePWReset)
	assert.Equal(t, false, resp.Users[0].OrganizationAdmin)
	assert.Equal(t, "2021-11-02 21:27:39", resp.Users[0].CreateDate)
	assert.Equal(t, "testington@testers.com", resp.Users[0].EmailAddress)
	assert.Equal(t, "testuser1", resp.Users[0].Username)
	assert.Equal(t, "Default Organization", resp.Users[0].OrganizationName)
	assert.Equal(t, false, resp.Users[0].ConsoleAccessDenied)
	assert.Equal(t, 2, resp.Users[0].UserID)
	assert.Equal(t, false, resp.Users[0].DomainAdmin)
	assert.Equal(t, true, resp.Users[0].ActiveApiKeyPresent)
	assert.Equal(t, "divvyuser:2:", resp.Users[0].ResourceID)
	assert.Equal(t, 0, resp.Users[0].OwnedResources)
	assert.Equal(t, 0, resp.Users[0].ConsecutiveFailedLoginAttempts)
	assert.Equal(t, false, resp.Users[0].Suspended)
	assert.Equal(t, "2022-04-11 15:23:19", resp.Users[0].LastLoginTime)
	assert.Equal(t, "Timmy Testington", resp.Users[0].Name)
	assert.Equal(t, false, resp.Users[1].OrganizationAdmin)
	assert.Equal(t, false, resp.Users[1].DomainAdmin)
	assert.Equal(t, "Default Organization", resp.Users[1].OrganizationName)
	assert.Equal(t, true, resp.Users[1].TwoFactorEnabled)
	assert.Equal(t, true, resp.Users[1].TwoFactorRequired)
	assert.Equal(t, false, resp.Users[1].ActiveApiKeyPresent)
	assert.Equal(t, "2022-04-11 15:11:34", resp.Users[1].CreateDate)
	assert.Equal(t, "divvyuser:9:", resp.Users[1].ResourceID)
	assert.Equal(t, 9, resp.Users[1].UserID)
	assert.Equal(t, false, resp.Users[1].DomainViewer)
	assert.Equal(t, 1, resp.Users[1].OrganizationID)
	assert.Equal(t, 0, resp.Users[1].ConsecutiveFailedLoginAttempts)
	assert.Equal(t, false, resp.Users[1].Suspended)
	assert.Equal(t, make([]string, 0), resp.Users[1].NavigationBlacklist)
	assert.Equal(t, "bbobb@bingo.xyz", resp.Users[1].EmailAddress)
	assert.Equal(t, "bbobb", resp.Users[1].Username)
	assert.Equal(t, "Billy Bobb", resp.Users[1].Name)
	assert.Equal(t, 0, resp.Users[1].Groups)
	assert.Equal(t, 0, resp.Users[1].OwnedResources)
	assert.Equal(t, true, resp.Users[1].RequirePWReset)
	assert.Equal(t, false, resp.Users[1].ConsoleAccessDenied)
	assert.Equal(t, "2022-04-11 15:12:34", resp.Users[2].CreateDate)
	assert.Equal(t, false, resp.Users[2].DomainAdmin)
	assert.Equal(t, false, resp.Users[2].RequirePWReset)
	assert.Equal(t, false, resp.Users[2].TwoFactorRequired)
	assert.Equal(t, "mjacks", resp.Users[2].Username)
	assert.Equal(t, false, resp.Users[2].TwoFactorEnabled)
	assert.Equal(t, "Default Organization", resp.Users[2].OrganizationName)
	assert.Equal(t, 1, resp.Users[2].OrganizationID)
	assert.Equal(t, 0, resp.Users[2].Groups)
	assert.Equal(t, 0, resp.Users[2].ConsecutiveFailedLoginAttempts)
	assert.Equal(t, false, resp.Users[2].ConsoleAccessDenied)
	assert.Equal(t, false, resp.Users[2].ActiveApiKeyPresent)
	assert.Equal(t, false, resp.Users[2].DomainViewer)
	assert.Equal(t, "mjacks@xyz.org", resp.Users[2].EmailAddress)
	assert.Equal(t, 10, resp.Users[2].UserID)
	assert.Equal(t, true, resp.Users[2].OrganizationAdmin)
	assert.Equal(t, 5, resp.Users[2].OwnedResources)
	assert.Equal(t, false, resp.Users[2].Suspended)
	assert.Equal(t, make([]string, 0), resp.Users[2].NavigationBlacklist)
	assert.Equal(t, "Okta", resp.Users[2].ServerName)
	assert.Equal(t, "divvyuser:10:", resp.Users[2].ResourceID)
	assert.Equal(t, "Mitchell Jacks", resp.Users[2].Name)
	assert.Equal(t, 3, resp.TotalCount)
	teardown()
}

func TestUsers_CreateUser(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_user_response.json"))
	})
	mux.HandleFunc("/v2/public/users/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_user_details_response.json"))
	})
	mux.HandleFunc("/v2/prototype/domains/admins/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_user_admin_details_response.json"))
	})

	resp, err := client.CreateUser(LocalUser{
		Name:              "Boaty McBoatFace",
		EmailAddress:      "boat@boatface.com",
		Username:          "Boatface",
		AccessLevel:       "BASIC_USER",
		TwoFactorRequired: false,
	})
	assert.NoError(t, err)
	assert.Equal(t, false, resp.DomainViewer)
	assert.Equal(t, "boat@boatface.com", resp.EmailAddress)
	assert.Equal(t, 1, resp.OrganizationID)
	assert.Equal(t, 3600, resp.SessionTimeoutSeconds)
	assert.Equal(t, true, resp.RequirePWReset)
	assert.Equal(t, "divvyuser:4:", resp.ResourceID)
	assert.Equal(t, "Boaty McBoatFace", resp.Name)
	assert.Equal(t, 4, resp.UserID)
	assert.Equal(t, true, resp.TwoFactorEnabled)
	assert.Equal(t, "local", resp.Auth)
	assert.Equal(t, "2023/01/01, 01:01:01 UTC", resp.SessionExpiration)
	assert.Equal(t, 1, resp.AuthenticationServerID)
	assert.Equal(t, false, resp.OrganizationAdmin)
	assert.Equal(t, false, resp.DomainAdmin)
	assert.Equal(t, "Default Organization", resp.OrganizationName)
	assert.Equal(t, false, resp.TwoFactorRequired)
	assert.Equal(t, 3600, resp.SessionTTL)
	assert.Equal(t, "2022-01-02 12:34:56", resp.CreateDate)
	assert.Equal(t, "1a2bcd3e-45fg-67hi-jklm-8o9p1q2r3st4u", resp.AWSDefaultExternalID)
	assert.Equal(t, false, resp.ApiKeyGenerationAllowed)
	assert.Equal(t, "Boatface", resp.Username)
	assert.Equal(t, false, resp.AuthPluginExists)
	assert.Equal(t, make([]string, 0), resp.NavigationBlacklist)
	assert.Equal(t, "divvy-light-theme", resp.Theme)
	assert.Equal(t, "", resp.Rapid7OrgID)
}

func TestUsers_CreateUser_NoName(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/user_list_response.json"))
	})

	_, err := client.CreateUser(LocalUser{
		EmailAddress:      "boat@boatface.com",
		Username:          "Boatface",
		AccessLevel:       "BASIC_USER",
		TwoFactorRequired: false,
	})
	assert.Error(t, err)
	teardown()
}

func TestUsers_CreateUser_NoEmail(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/user_list_response.json"))
	})

	_, err := client.CreateUser(LocalUser{
		Name:              "Boaty McBoatface",
		Username:          "Boatface",
		AccessLevel:       "BASIC_USER",
		TwoFactorRequired: false,
	})
	assert.Error(t, err)
	teardown()
}

func TestUsers_CreateUser_NoUsername(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/user_list_response.json"))
	})

	_, err := client.CreateUser(LocalUser{
		Name:              "Boaty McBoatface",
		EmailAddress:      "boat@boatface.com",
		AccessLevel:       "BASIC_USER",
		TwoFactorRequired: false,
	})
	assert.Error(t, err)
	teardown()
}

func TestUsers_CreateUser_BadAccessLevel(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/users/create", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/user_list_response.json"))
	})

	_, err := client.CreateUser(LocalUser{
		Name:              "Boaty McBoatface",
		EmailAddress:      "boat@boatface.com",
		AccessLevel:       "GOD_MODE",
		TwoFactorRequired: false,
	})
	assert.Error(t, err)
	teardown()
}

func TestUsers_CreateAPIUser_NoName(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/create_api_only_user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_api_user_response.json"))
	})

	expire_date := time.Now().Add(720 * time.Hour)

	_, err := client.CreateAPIUser(APIUser{
		EmailAddress:   "api@mcboatface.com",
		Username:       "api_boatface",
		ExpirationDate: expire_date.Unix(),
	})
	assert.Error(t, err)
	teardown()
}

func TestUsers_CreateAPIUser_NoEmail(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/create_api_only_user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_api_user_response.json"))
	})

	expire_date := time.Now().Add(720 * time.Hour)

	_, err := client.CreateAPIUser(APIUser{
		Name:           "API McBoatface",
		Username:       "api_boatface",
		ExpirationDate: expire_date.Unix(),
	})
	assert.Error(t, err)
	teardown()
}

func TestUsers_CreateAPIUser_NoUsername(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/create_api_only_user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_api_user_response.json"))
	})

	expire_date := time.Now().Add(720 * time.Hour)

	_, err := client.CreateAPIUser(APIUser{
		Name:           "API McBoatface",
		EmailAddress:   "api@mcboatface.com",
		ExpirationDate: expire_date.Unix(),
	})
	assert.Error(t, err)
	teardown()
}

func TestUsers_CreateAPIUser_NoExpirationDate(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/create_api_only_user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_api_user_response.json"))
	})

	mux.HandleFunc("/v2/public/users/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_user_details_response.json"))
	})
	mux.HandleFunc("/v2/prototype/domains/admins/list", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/create_user_admin_details_response.json"))
	})

	_, err := client.CreateAPIUser(APIUser{
		Name:         "API McBoatface",
		EmailAddress: "api@mcboatface.com",
		Username:     "api_boatface",
	})
	assert.NoError(t, err)
	teardown()
}
