package insightcloudsecClient

type Group struct {
	ID                    int    `json:"group_id"`
	ResourceID            string `json:"resource_id"`
	Name                  string `json:"name"`
	Users                 int    `json:"users"`
	Roles                 int    `json:"roles"`
	EntitlementConfigured bool   `json:"entitlement_configured"`
}

type Groups struct {
	Groups []Group `json:"groups"`
}
