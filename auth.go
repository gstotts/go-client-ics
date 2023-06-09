package insightcloudsecClient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthStruct
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse
type AuthResponse struct {
	UserID           int    `json:"user_id"`
	Name             string `json:"user_name"`
	Email            string `json:"user_email"`
	SessionID        string `json:"session_id"`
	Timeout          int    `json:"session_timeout"`
	DomainAdmin      bool   `json:"domain_admin"`
	CustomerID       string `json:"customer_id"`
	DomainViewer     bool   `json:"domain_viewer"`
	AuthPluginExists bool   `json:"auth_plugin_exist"`
}

// ApikeyRequest
type ApikeyRequest struct {
	UserID    string `json:"user_id"`
	KeyLength int32  `json:"key_length"`
}

// ApikeyResponse
type ApikeyResponse struct {
	APIKey string `json:"apikey"`
}

// Login to InsightCloudSec
func (c *Client) Login() (AuthResponse, error) {

	// Verify AuthStruct is not blank
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return AuthResponse{}, fmt.Errorf("missing username and/or password")
	}

	// Make login request
	body, err := c.makeRequest(http.MethodPost, "/v2/public/user/login", c.Auth)
	if err != nil {
		return AuthResponse{}, err
	}

	// Unmarshal Data
	resp := AuthResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return AuthResponse{}, err
	}

	return resp, nil
}

func (c *Client) CreateAPIKey(key_length int) (string, error) {

	// Create Payload
	data := ApikeyRequest{
		UserID:    fmt.Sprintf("%d", c.UserDetails.UserID),
		KeyLength: int32(key_length),
	}

	// Make Request
	body, err := c.makeRequest(http.MethodPost, "/v2/public/apikey/create", data)
	if err != nil {
		return "", err
	}

	// Unmarshal Response
	resp := ApikeyResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}

	return resp.APIKey, nil
}
