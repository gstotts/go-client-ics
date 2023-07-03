package insightcloudsecClient

import "net/http"

// WhitelabelImage
type WhitelabelImage struct {
	Image string `json:"image"`
}

func (c *Client) SetCustomLogo(image_url string) error {
	// Sets the logo to a custom image png.  Can be base64-encoded png url.  Approx 450 x 115 px

	return c.makeRequest(http.MethodPost, "/v2/whitelabel/set-custom-logo", WhitelabelImage{image_url}, nil)
}
