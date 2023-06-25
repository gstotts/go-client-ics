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

	testCases := []struct {
		resource_id     string
		name            string
		user_id         int
		org_admin       bool
		domain_admin    bool
		domain_view     bool
		email           string
		username        string
		org_name        string
		org_id          int
		tfa_enabled     bool
		tfa_req         bool
		groups          int
		owned_resources int
		failed_logins   int
		suspended       bool
		login_time      string
		nav_blacklist   []string
		server_name     string
		pw_reset        bool
		ca_denied       bool
		api_key         bool
		create          string
	}{
		{"divvyuser:2:", "Timmy Testington", 2, false, false, false, "testington@testers.com", "testuser1", "Default Organization", 1, false, false, 1, 0, 0, false, "2022-04-11 15:23:19", []string{}, "", false, false, true, "2021-11-02 21:27:39"},
		{"divvyuser:9:", "Billy Bobb", 9, false, false, false, "bbobb@bingo.xyz", "bbobb", "Default Organization", 1, true, true, 0, 0, 0, false, "", []string{}, "", true, false, false, "2022-04-11 15:11:34"},
		{"divvyuser:10:", "Mitchell Jacks", 10, true, false, false, "mjacks@xyz.org", "mjacks", "Default Organization", 1, false, false, 0, 5, 0, false, "", []string{}, "Okta", false, false, false, "2022-04-11 15:12:34"},
	}
	resp, err := client.ListBasicUsers()
	assert.NoError(t, err)
	assert.Equal(t, 3, resp.TotalCount)

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("User Test: %s", tc.resource_id), func(t *testing.T) {
			assert.Equal(t, tc.resource_id, resp.Users[i].ResourceID)
			assert.Equal(t, tc.name, resp.Users[i].Name)
			assert.Equal(t, tc.user_id, resp.Users[i].UserID)
			assert.Equal(t, tc.org_admin, resp.Users[i].OrganizationAdmin)
			assert.Equal(t, tc.domain_admin, resp.Users[i].DomainAdmin)
			assert.Equal(t, tc.domain_view, resp.Users[i].DomainViewer)
			assert.Equal(t, tc.email, resp.Users[i].EmailAddress)
			assert.Equal(t, tc.username, resp.Users[i].Username)
			assert.Equal(t, tc.org_name, resp.Users[i].OrganizationName)
			assert.Equal(t, tc.org_id, resp.Users[i].OrganizationID)
			assert.Equal(t, tc.tfa_enabled, resp.Users[i].TwoFactorEnabled)
			assert.Equal(t, tc.tfa_req, resp.Users[i].TwoFactorRequired)
			assert.Equal(t, tc.groups, resp.Users[i].Groups)
			assert.Equal(t, tc.owned_resources, resp.Users[i].OwnedResources)
			assert.Equal(t, tc.failed_logins, resp.Users[i].ConsecutiveFailedLoginAttempts)
			assert.Equal(t, tc.suspended, resp.Users[i].Suspended)
			assert.Equal(t, tc.login_time, resp.Users[i].LastLoginTime)
			assert.Equal(t, tc.nav_blacklist, resp.Users[i].NavigationBlacklist)
			assert.Equal(t, tc.pw_reset, resp.Users[i].RequirePWReset)
			assert.Equal(t, tc.ca_denied, resp.Users[i].ConsoleAccessDenied)
			assert.Equal(t, tc.api_key, resp.Users[i].ActiveApiKeyPresent)
			assert.Equal(t, tc.create, resp.Users[i].CreateDate)
			assert.Equal(t, tc.server_name, resp.Users[i].ServerName)

		})
	}
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

	testCases := []struct {
		name      string
		email     string
		username  string
		access    string
		tfa       bool
		test_name string
		err       bool
	}{
		{"Boaty McBoatFace", "boat@boatface.com", "Boatface", "BASIC_USER", false, "Valid User", false},
		{"", "boat@boatface.com", "Boatface", "BASIC_USER", false, "No Username", true},
		{"Boaty McBoatFace", "", "Boatface", "BASIC_USER", false, "No Email", true},
		{"Boaty McBoatFace", "boat@boatface.com", "", "BASIC_USER", false, "No Username", true},
		{"Boaty McBoatFace", "boat@boatface.com", "Boatface", "GOD_MODE", false, "Invalid Access Setting", true},
	}

	for _, tc := range testCases {
		t.Run(tc.test_name, func(t *testing.T) {
			resp, err := client.CreateUser(LocalUser{
				Name:              tc.name,
				EmailAddress:      tc.email,
				Username:          tc.username,
				AccessLevel:       tc.access,
				TwoFactorRequired: tc.tfa,
			})
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tc.access == "BASIC_USER" {
					assert.False(t, resp.OrganizationAdmin)
					assert.False(t, resp.DomainAdmin)
					assert.False(t, resp.DomainViewer)
				}
				if tc.access == "ORGANIZATION_ADMIN" {
					assert.True(t, resp.OrganizationAdmin)
					assert.False(t, resp.DomainAdmin)
					assert.False(t, resp.DomainViewer)
				}
				if tc.access == "DOMAIN_VIEWER" {
					assert.False(t, resp.OrganizationAdmin)
					assert.False(t, resp.DomainAdmin)
					assert.True(t, resp.DomainViewer)
				}
				if tc.access == "DOMAIN_ADMIN" {
					assert.False(t, resp.OrganizationAdmin)
					assert.True(t, resp.DomainAdmin)
					assert.False(t, resp.DomainViewer)
				}
				assert.Equal(t, tc.name, resp.Name)
				assert.Equal(t, tc.email, resp.EmailAddress)
				assert.Equal(t, tc.username, resp.Username)
				assert.Equal(t, tc.tfa, resp.TwoFactorRequired)
				assert.Equal(t, "divvyuser:4:", resp.ResourceID)
				assert.Equal(t, 1, resp.OrganizationID)
				assert.Equal(t, 3600, resp.SessionTimeoutSeconds)
				assert.Equal(t, true, resp.RequirePWReset)
				assert.Equal(t, 4, resp.UserID)
				assert.Equal(t, true, resp.TwoFactorEnabled)
				assert.Equal(t, "local", resp.Auth)
				assert.Equal(t, "2023/01/01, 01:01:01 UTC", resp.SessionExpiration)
				assert.Equal(t, 1, resp.AuthenticationServerID)
				assert.Equal(t, "Default Organization", resp.OrganizationName)
				assert.Equal(t, false, resp.TwoFactorRequired)
				assert.Equal(t, 3600, resp.SessionTTL)
				assert.Equal(t, "2022-01-02 12:34:56", resp.CreateDate)
				assert.Equal(t, "1a2bcd3e-45fg-67hi-jklm-8o9p1q2r3st4u", resp.AWSDefaultExternalID)
				assert.Equal(t, false, resp.ApiKeyGenerationAllowed)
				assert.Equal(t, false, resp.AuthPluginExists)
				assert.Equal(t, make([]string, 0), resp.NavigationBlacklist)
				assert.Equal(t, "divvy-light-theme", resp.Theme)
				assert.Equal(t, "", resp.Rapid7OrgID)
				assert.Equal(t, "..u^0bVZ%,#a", resp.TemporaryPW)
				assert.Equal(t, "2023-06-10T01:35:39Z", resp.TempPWExpiration)
			}
		})
	}
	teardown()
}

func TestUsers_CreateAPIUser(t *testing.T) {
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

	expire_date := time.Now().Add(720 * time.Hour)

	resp, err := client.CreateAPIUser(APIUser{
		Name:           "API McBoatface",
		EmailAddress:   "api@mcboatface.com",
		Username:       "api_boatface",
		ExpirationDate: expire_date.Unix(),
	})
	assert.NoError(t, err)
	assert.Equal(t, "API McBoatface", resp.Name)
	assert.Equal(t, "api@mcboatface.com", resp.EmailAddress)
	assert.Equal(t, "api_boatface", resp.Username)
	assert.Equal(t, "n8kP6d2SaWw7kVNgMWuUd3wN1cTddqW2aKWLpg18Lv-h2ceMymg", resp.ApiKey)
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

func TestUsers_DeleteUser(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/prototype/user/divvyuser:5:/delete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	assert.NoError(t, client.DeleteUser("divvyuser:5:"))
	teardown()
}

func TestUsers_ConvertUserToAPIUser(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/update_to_api_only_user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/convert_to_api_user_response.json"))
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

	resp, err := client.ConvertUserToAPIUser(8)
	assert.NoError(t, err)
	assert.Equal(t, 8, resp.UserID)
	assert.Equal(t, "E009_o_beBcNI8Rdp3si_KTmL38c_MFd08tUWpSMkREUDCVCaqo", resp.ApiKey)
	teardown()
}

func TestUsers_2FA_Status(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/tfa_state", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/user_tfa_status_response.json"))
	})
	resp, err := client.Get2FAStatus(8)
	assert.NoError(t, err)
	assert.Equal(t, true, resp.Enabled)
	assert.Equal(t, false, resp.Required)
	teardown()
}

func TestUsers_Enable2FA(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/tfa_enable", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getJSONFile("users/enable_2fa_response.json"))
	})
	resp, err := client.Enable2FA()
	assert.NoError(t, err)
	assert.Equal(t, "AABB2CDEMFGGAB34C", resp.Secret)
	teardown()
}

func TestUsers_Disable2FA(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/tfa_disable", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	err := client.Disable2FA(999)
	assert.NoError(t, err)
	teardown()
}

func TestUsers_ConsoleAccessDeniedFlag(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/user/update_console_access", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	err := client.UpdateConsoleAccessDeniedFlag(999, false)
	assert.NoError(t, err)
	teardown()
}

func TestUsers_DeactivateAPIKeys(t *testing.T) {
	setup()
	mux.HandleFunc("/v2/public/apikey/deactivate", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
	err := client.DeactivateAPIKeys(999)
	assert.NoError(t, err)
	teardown()
}
