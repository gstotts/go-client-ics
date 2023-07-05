package insightcloudsecClient

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func (c *Client) ListCloudTypes() (CloudTypes, error) {
	// Returns a list of all cloud types

	var cloudTypes CloudTypes
	err := c.makeRequest(http.MethodGet, "/v2/public/cloudtypes/list", nil, &cloudTypes)
	return cloudTypes, err
}

func (c *Client) ListClouds() ([]Cloud, error) {
	// List availble configured clouds

	var list Clouds
	err := c.makeRequest(http.MethodGet, "/v2/public/clouds/list", nil, &list)
	return list.Clouds, err
}

func (c *Client) GetCloudByID(id any) (Cloud, error) {
	// Returns a specific cloud of given id

	list, err := c.ListClouds()
	if err != nil {
		return Cloud{}, err
	}

	var toCompare string
	switch v := id.(type) {
	case int:
		toCompare = fmt.Sprintf("divvyorganizationservice:%d", v)
	case string:
		r, _ := regexp.Compile(`divvyorganizationservice:\d+`)
		if r.MatchString(v) {
			toCompare = v
		} else {
			return c.GetCloudByName(v)
		}
	default:
		return Cloud{}, fmt.Errorf("id must be of type string or int, got %T", v)
	}

	var desiredCloud Cloud
	for _, cloud := range list {
		if cloud.ResourceID == toCompare {
			desiredCloud = cloud
			desiredCloud.client = c
			return desiredCloud, nil
		}
	}

	return Cloud{}, fmt.Errorf("unable to find cloud of id: %d", id)
}

func (c *Client) GetCloudByName(name string) (Cloud, error) {
	// Return a cloud of given name

	list, err := c.ListClouds()
	if err != nil {
		return Cloud{}, err
	}

	var desiredCloud Cloud
	for _, cloud := range list {
		if strings.EqualFold(cloud.Name, name) {
			desiredCloud = cloud
			desiredCloud.client = c
			return desiredCloud, nil
		}
	}

	return Cloud{}, fmt.Errorf("unable to find cloud of name: %s", name)
}

func (c *Cloud) Update_Name(new_name string) error {
	// Updates the name of the cloud

	err := c.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/prototype/resource/%s/name/set", c.ResourceID), map[string]string{"name": new_name}, nil)
	if err != nil {
		return err
	}
	*c, _ = c.client.GetCloudByID(c.ResourceID)
	return nil
	// return c.client.GetCloudByID(c.ResourceID)
}

func (c *Cloud) Update(name string, account_number int, authentication_type, cloud_type, session_name, external_id string) {
}
