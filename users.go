package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"strconv"
)

func (c *Client) CurrentUserInfo() (User, error) {
	// Returns the current user information

	resp := User{}
	err := c.makeRequest(http.MethodGet, "/v2/public/user/info", nil, &resp)

	return resp, err
}

func (c *Client) getUsers(method, url string) (Users, error) {
	// Makes GET to Users API and Returns User List

	resp := Users{}
	err := c.makeRequest(method, url, nil, &resp)

	return resp, err
}

func (c *Client) ListBasicUsers() (Users, error) {
	// Lists All Standard Users (non-Domain Admins)

	users, err := c.getUsers(http.MethodGet, "/v2/public/users/list")
	return users, err
}

func (c *Client) ListAdmins() (Users, error) {
	// List Admin Users

	users, err := c.getUsers(http.MethodPost, "/v2/prototype/domains/admins/list")
	users.TotalCount = len(users.Users)
	return users, err
}

func (c *Client) ListUsers() (Users, error) {
	// List All Users - Basic and Admin

	basic_users, err := c.ListBasicUsers()
	if err != nil {
		return Users{TotalCount: 0, Users: []User{}}, nil
	}

	admins, err := c.ListAdmins()
	if err != nil {
		return Users{TotalCount: 0, Users: []User{}}, nil
	}

	var combined Users
	combined.TotalCount = basic_users.TotalCount + admins.TotalCount
	combined.Users = append(basic_users.Users, admins.Users...)

	return combined, nil
}

func (c *Client) GetUserByUsername(username string) (User, error) {
	// Returns the user of the given username

	all_users, err := c.ListUsers()
	if err != nil {
		return User{}, err
	}

	for i, user := range all_users.Users {
		if user.Username == username {
			return all_users.Users[i], nil
		}
	}

	return User{}, fmt.Errorf("unable to find username %s", username)
}

func (c *Client) GetUserByID(user_id int) (User, error) {
	// Returns the user of the given user_id

	all_users, err := c.ListUsers()
	if err != nil {
		return User{}, err
	}

	for i, user := range all_users.Users {
		if user.UserID == user_id {
			return all_users.Users[i], nil
		}
	}

	return User{}, fmt.Errorf("unable to find user of user_id %d", user_id)
}

func (c *Client) postUser(url string, data interface{}) (userTempPasswordResponse, error) {
	// Makes POST to Users API and returns User

	resp := userTempPasswordResponse{}
	err := c.makeRequest(http.MethodPost, url, data, &resp)

	return resp, err
}

func (c *Client) CreateUser(user LocalUser) (User, error) {
	// Creates an InsightCloudSec User account

	if user.Name == "" || user.AccessLevel == "" || user.EmailAddress == "" || user.Username == "" {
		return User{}, fmt.Errorf("must set user's name, emailaddress, username and accesslevel")
	}

	if !isValidAccessLevel(user.AccessLevel) {
		return User{}, fmt.Errorf("accesslevel must be one of: BASIC_USER, ORGANIZATION_ADMIN, DOMAIN_VIEWER, or DOMAIN_ADMIN")
	}

	// Make Request
	user_resp, err := c.postUser("/v2/public/user/create", userCreateRequest(user))
	if err != nil {
		return User{}, err
	}

	// Get Full User Details
	details, err := c.GetUserByID(user_resp.UserID)
	// Append Full User Details with the Temporary PW
	details.TempPWExpiration = user_resp.TempPWExpiration
	details.TemporaryPW = user_resp.TemporaryPW

	return details, err
}

func (c *Client) CreateAPIUser(user APIUser) (User, error) {
	// Creates an InsightCloudSec API Only User account

	if user.Name == "" || user.EmailAddress == "" || user.Username == "" {
		return User{}, fmt.Errorf("must set api users's name, emailaddress, and username")
	}

	resp := userCreateAPIKeyResponse{}
	err := c.makeRequest(http.MethodPost, "/v2/public/user/create_api_only_user", user, &resp)
	if err != nil {
		return User{}, err
	}

	// Get Full API User Details
	details, err := c.GetUserByID(resp.UserID)
	details.ApiKey = resp.ApiKey

	return details, err
}

// NEED AUTH SERVERS SETUP PRIOR
//
// func (c *Client) CreateSAMLUser(user SAMLUser) (User, error) {
// 	resp := User{}
// 	return resp, nil
// }

func (c *Client) DeleteUser(resource_id string) error {
	// Deletes an InsightCloudSec user of given resource ID

	return c.makeRequest(http.MethodDelete, fmt.Sprintf("/v2/prototype/user/%s/delete", resource_id), nil, nil)
}

func (c *Client) Get2FAStatus(user_id int) (MFAStatus, error) {
	// Returns whether 2FA is enabled and/or required for the user of given ID

	// Make Request
	data := map[string]int{"user_id": user_id}
	resp := MFAStatus{}
	err := c.makeRequest(http.MethodPost, "/v2/public/user/tfa_state", data, &resp)

	return resp, err
}

func (c *Client) Enable2FA() (OTPSecret, error) {
	// Enables 2FA for current user and returns OTP Secret to utilize

	resp := OTPSecret{}
	err := c.makeRequest(http.MethodPost, "/v2/public/user/tfa_enable", nil, &resp)
	return resp, err
}

func (c *Client) Disable2FA(user_id int) error {
	// Disables 2FA for the user of given ID

	return c.makeRequest(http.MethodPost, "/v2/public/user/tfa_disable", map[string]int{"user_id": user_id}, nil)
}

func (c *Client) ConvertUserToAPIUser(user_id int) (User, error) {
	// Converts a normal user to an api-only user

	data := map[string]string{"user_id": strconv.Itoa(user_id)}
	resp := userConvertToAPIUserResponse{}
	err := c.makeRequest(http.MethodPost, "/v2/public/user/update_to_api_only_user", data, &resp)
	if err != nil {
		return User{}, err
	}

	// Get Full API User Details
	details, err := c.GetUserByID(user_id)
	details.ApiKey = resp.ApiKey
	return details, err
}

func (c *Client) UpdateConsoleAccessDeniedFlag(user_id int, console_access_denied bool) error {
	// Sets the console access for the given user of user_id
	return c.makeRequest(http.MethodPost, "/v2/public/user/update_console_access", map[string]interface{}{"user_id": strconv.Itoa(user_id), "console_access_denied": console_access_denied}, nil)
}

func (c *Client) DeactivateAPIKeys(user_id int) error {
	// Deactivates API Keys for a given user of user_id

	return c.makeRequest(http.MethodPost, "/v2/public/apikey/deactivate", map[string]string{"user_id": strconv.Itoa(user_id)}, nil)
}

func isValidAccessLevel(level string) bool {
	switch level {
	case "BASIC_USER":
		return true
	case "ORGANIZATION_ADMIN":
		return true
	case "DOMAIN_VIEWER":
		return true
	case "DOMAIN_ADMIN":
		return true
	default:
		return false
	}
}
