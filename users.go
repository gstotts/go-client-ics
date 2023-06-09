package insightcloudsecClient

import (
	"encoding/json"
	"net/http"
)

func (c *Client) CurrentUserInfo() (UserInfoResponse, error) {

	// Make Request
	body, err := c.makeRequest(http.MethodGet, "/v2/public/user/info", "")
	if err != nil {
		return UserInfoResponse{}, err
	}

	// Unmarshal Response
	resp := UserInfoResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return UserInfoResponse{}, err
	}

	return resp, nil
}

func (c *Client) ListUsers() (UserListResponse, error) {

	// Make Request
	body, err := c.makeRequest(http.MethodGet, "/v2/public/users/list", nil)
	if err != nil {
		return UserListResponse{TotalCount: 0, Users: []UserListItem{}}, err
	}

	// Unmarshal Response
	resp := UserListResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return UserListResponse{TotalCount: 0, Users: []UserListItem{}}, err
	}

	return resp, nil
}
