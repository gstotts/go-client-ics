package insightcloudsecClient

type Role struct {
	ResourceID          string   `json:"resource_id,omitempty"`
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
	CloudScopes         []string `json:"cloud_scopes,omitempty"`
	ResourceGroupScopes []string `json:"resource_group_scopes,omitempty"`
	BadgeScopes         []string `json:"badge_scopes,omitempty"`
	Groups              []string `json:"groups,omitempty"`
}

type Roles struct {
	Roles []Role `json:"roles"`
}

type rolesUpdateScopeRequest struct {
	ResourceIDs           []string `json:"resource_ids"`
	DeprecatedResourceIDs []string `json:"deprecated_resource_ids"`
}

type rolesUpdateUserGroupsRequest struct {
	GroupResourceIDs []string `json:"group_resource_ids"`
}
