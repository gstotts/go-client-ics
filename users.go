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

func (c *Client) ListUsers() (UserList, error) {
	// Lists All Standard Users (non-Domain Admins)

	// Make Request
	body, err := c.makeRequest(http.MethodGet, "/v2/public/users/list", nil)
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

func (c *Client) CreateUser(user LocalUser) (User, error) {
	// Creates an InsightCloudSec User account

	if user.Name == "" || user.AccessLevel == "" || user.EmailAddress == "" || user.Password == "" || user.Username == "" {
		return User{}, fmt.Errorf("must set user's name, emailaddress, password, username and accesslevel")
	}

	if !isValidAccessLevel(user.AccessLevel) {
		return User{}, fmt.Errorf("accesslevel must be one of: BASIC_USER, ORGANIZATION_ADMIN, DOMAIN_VIEWER, or DOMAIN_ADMIN")
	}

	// Make Request
	body, err := c.makeRequest(http.MethodPost, "/v2/public/user/create", userCreateRequest{
		Name:              user.Name,
		EmailAddress:      user.EmailAddress,
		Username:          user.Username,
		AccessLevel:       user.AccessLevel,
		Password:          user.Password,
		TwoFactorRequired: user.TwoFactorRequired,
		ConfirmPassword:   user.Password,
	})
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

func (c *Client) CreateAPIUser(user APIUser) (User, error) {
	resp := User{}
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
