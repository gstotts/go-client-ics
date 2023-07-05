package insightcloudsecClient

type CloudType struct {
	ID     string `json:"cloud_type_id"`
	Access string `json:"cloud_access"`
	Name   string `json:"name"`
}

type CloudTypes struct {
	CloudTypes []CloudType `json:"clouds"`
}
