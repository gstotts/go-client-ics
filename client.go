package insightcloudsecClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Default URL for Local Testing
const HostURL string = "http://localhost:8001"

// Client
type Client struct {
	HostURL     string
	HTTPClient  *http.Client
	Token       string
	APIKey      string
	Auth        AuthStruct
	UserDetails AuthResponse
}

// NewClient
func NewClient(host, username, password, apikey *string) (*Client, error) {

	// Create basic client
	c := Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		HostURL:    HostURL,
	}

	// Set host if provided
	if host != nil {
		c.HostURL = *host
	}

	// Set apikey if provided
	if apikey != nil {
		c.APIKey = *apikey
	}

	// If username and password are not set, return client as is
	if username == nil || password == nil {
		return &c, nil
	}

	// Create AuthStruct for Login
	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	// Login with Username and Password to receive Session ID
	resp, err := c.Login()
	if err != nil {
		return nil, err
	}

	// Set sessionID token
	c.Token = resp.SessionID
	c.UserDetails = resp

	// Return the Client and no error
	return &c, nil
}

func (c *Client) makeRequest(method, path string, data interface{}, result any) error {

	// Get Raw Request Response
	resp, err := c.makeRawRequest(method, path, data)
	if err != nil {
		return err
	}

	// Read Response Data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Validate Status Code
	if resp.StatusCode != http.StatusOK {
		var api_error APIErrorResponse
		_ = json.Unmarshal(body, &api_error)
		return fmt.Errorf("\n      HTTP Status: %d,\n   API Error Type: %s,\nAPI Error Message: %s", resp.StatusCode, api_error.ErrorType, api_error)
	}

	fmt.Println(body, path)
	if result != nil {
		err = json.Unmarshal(body, &result)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) makeRawRequest(method, path string, data interface{}) (*http.Response, error) {

	// Marshall Data for Payload
	byte_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Build Request
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.HostURL, path), bytes.NewBuffer(byte_data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", c.Token)
	req.Header.Set("API-Key", c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "insightcloudsec-client-go")

	// Get Response from InsightCloudSec API
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
