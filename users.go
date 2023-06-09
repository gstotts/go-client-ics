package insightcloudsecClient

import (
	"encoding/json"
	"net/http"
)

func (c *Client) CurrentUserInfo() (User, error) {
	// Returns the current user information

	// Make Request
	body, err := c.makeRequest(http.MethodGet, "/v2/public/user/info", "")
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

func (c *Client) CreateUser(user NewUser) (User, error) {
	resp := User{}
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
