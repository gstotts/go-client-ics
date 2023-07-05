package insightcloudsecClient

type CloudTypes struct {
	Clouds []CloudType `json:"clouds"`
}

type CloudType struct {
	ID     string `json:"cloud_type_id"`
	Access string `json:"cloud_access"`
	Name   string `json:"name"`
}
