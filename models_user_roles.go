package insightcloudsecClient

type Role struct {
	ResourceID          string   `json:"resource_id"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	BadgeFilterOperator string   `json:"badge_filter_operator"`
	AllPermissions      bool     `json:"all_permissions"`
	View                bool     `json:"view"`
	Provision           bool     `json:"provision"`
	Manage              bool     `json:"manage"`
	Delete              bool     `json:"delete"`
	AddCloud            bool     `json:"add_cloud"`
	DeleteCloud         bool     `json:"delete_cloud"`
	GlobalScope         bool     `json:"global_scope"`
	CloudScopes         []string `json:"cloud_scopes"`
	ResourceGroupScopes []string `json:"resource_group_scopes"`
	BadgeScopes         []string `json:"badge_scopes"`
	Groups              []string `json:"groups"`
}

type Roles struct {
	Roles Role `json:"roles"`
}
