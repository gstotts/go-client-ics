package insightcloudsecClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) badgeModifyRequest(add_to_path string, ids []string, b []Badge) error {
	payload := struct {
		IDs []string `json:"target_resource_ids"`
		Badges
	}{
		IDs: ids,
		Badges: Badges{
			Badges: b,
		},
	}

	// API can return 204 Status when successful, so must handle the response differently than makeRequest()
	resp, err := c.makeRawRequest(http.MethodPost, fmt.Sprintf("/v2/public/badges/%s", add_to_path), payload)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var api_error APIErrorResponse
		_ = json.Unmarshal(body, &api_error)
		return fmt.Errorf("\n      HTTP Status: %d,\n   API Error Type: %s,\nAPI Error Message: %s", resp.StatusCode, api_error.ErrorType, api_error)
	}

	return nil
}

func (c *Client) CreateBadges(target_resource_ids []string, b []Badge) error {
	// Creates the badges and associates them with the supplied targets of given ids

	return c.badgeModifyRequest("create", target_resource_ids, b)
}

func (c *Client) UpdateCloudBadges(resource_id string, b []Badge) error {
	// Ovewrites all the badges associated with the resource_id with the given list

	return c.badgeModifyRequest(fmt.Sprintf("%s/update", resource_id), []string{resource_id}, b)
}

func (c *Client) DeleteBadges(target_resource_ids []string, b []Badge) error {
	// Deletes the badges and associates them with the supplied targets of given ids

	return c.badgeModifyRequest("delete", target_resource_ids, b)
}

func (c *Client) ListResourceBadges(resource_id string) ([]Badge, error) {
	// Lists badges associated with given resource_id

	var resp []Badge
	err := c.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/badges/%s/list", resource_id), nil, &resp)
	return resp, err
}

func (c *Client) ListCloudsWithBadges() ([]BadgedCloud, error) {
	// Retunrs a list of cloud accounts that are badged

	var resp []BadgedCloud
	err := c.makeRequest(http.MethodPost, "/v2/public/badge/clouds/list", nil, &resp)
	return resp, err
}

func (c *Client) ListResourceBadgeCount(resource_ids []string) ([]BadgeCount, error) {
	// Returns a list of badge counts for all resources

	var resp BadgeResourceCount
	err := c.makeRequest(http.MethodPost, "/v2/public/badges/count", struct {
		IDs []string `json:"resource_ids"`
	}{IDs: resource_ids}, &resp)

	return resp.ResourceCount, err
}
