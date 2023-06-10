package insightcloudsecClient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CurrentUserInfo() (User, error) {
	// Returns the current user information

	// Make Request
	body, err := c.makeRequest(http.MethodGet, "/v2/public/user/info", nil)
	if err != nil {
		return User{}, err
	}

	// Unmarshal Response
	resp := User{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return User{}, err
	}

	return resp, nil
}

func (c *Client) getUsers(url string) (UserList, error) {
	// Makes GET to Users API and Returns User List

	// Make Request
	body, err := c.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		return UserList{TotalCount: 0, Users: []User{}}, err
	}

	// Unmarshal Response
	resp := UserList{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return UserList{TotalCount: 0, Users: []User{}}, err
	}

	return resp, nil
}

func (c *Client) ListBasicUsers() (UserList, error) {
	// Lists All Standard Users (non-Domain Admins)

	users, err := c.getUsers("/v2/public/users/list")
	return users, err
}

func (c *Client) ListAdmins() (UserList, error) {
	// List Admin Users

	users, err := c.getUsers("/v2/prototype/domains/admins/list")
	users.TotalCount = len(users.Users)
	return users, err
}

func (c *Client) ListUsers() (UserList, error) {
	// List All Users - Basic and Admin

	basic_users, err := c.ListBasicUsers()
	if err != nil {
		return UserList{TotalCount: 0, Users: []User{}}, nil
	}

	admins, err := c.ListAdmins()
	if err != nil {
		return UserList{TotalCount: 0, Users: []User{}}, nil
	}

	var combined UserList
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

	// Make Request
	body, err := c.makeRequest(http.MethodPost, url, data)
	if err != nil {
		return userTempPasswordResponse{}, err
	}

	// Unmarshal Response
	resp := userTempPasswordResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return userTempPasswordResponse{}, err
	}

	return resp, nil
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

	// Make Request
	body, err := c.makeRequest(http.MethodPost, "/v2/public/user/create", user)
	if err != nil {
		return User{}, err
	}

	// Unmarshal Response
	resp := User{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return User{}, nil
	}

	return resp, nil
}

func (c *Client) CreateSAMLUser(user SAMLUser) (User, error) {
	resp := User{}
	return resp, nil
}

func (c *Client) DeleteUser(resource_id string) error {
	return nil
}

func (c *Client) Get2FAStatus(user_id int) (MFAStatus, error) {
	resp := MFAStatus{}
	return resp, nil
}

func (c *Client) Enable2FA() (OTPSecret, error) {
	resp := OTPSecret{}
	return resp, nil
}

func (c *Client) Disable2FA(user_id int) error {
	return nil
}

func (c *Client) ConvertUserToAPIUser(user_id int) (string, error) {
	key := ""
	return key, nil
}

func (c *Client) UpdateConsoleAccessDeniedFlag(user_id string, console_access_denied bool) error {
	return nil
}

func (c *Client) DeactivateAPIKeys(user_id string) error {
	return nil
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
