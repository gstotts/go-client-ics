package insightcloudsecClient

type PrincipalActivity struct {
	ExecutedActions []PrincipalActions `json:"executed_actions"`
	Metadata        Metadata           `json:"metadata"`
	Principal       Principal          `json:"principal"`
}

type Principal struct {
	Name                  string `json:"name"`
	AccountID             string `json:"account_id,omitempty"`
	CloudTypeID           string `json:"cloud_type_id,omitempty"`
	NamespaceID           string `json:"namespace_id,omitempty"`
	OrganizationID        int    `json:"organization_id,omitempty"`
	OrganizationServiceID int    `json:"organization_service_id,omitempty"`
	ProviderID            string `json:"provider_id,omitempty"`
	ResourceID            string `json:"resource_id"`
	ResourceType          string `json:"resource_type"`
}

type Metadata struct {
	Errors Errs `json:"errors"`
}

type Errs struct {
	Date   string `json:"date"`
	Reason string `json:"reason"`
}
type PrincipalActions struct {
	Name         string `json:"action"`
	Count        int    `json:"count"`
	LastExecuted string `json:"last_executed_date"`
}

type PrincipalPermissions struct {
	End          string                 `json:"end"`
	Page         int                    `json:"page"`
	Permissions  Permissions            `json:"permissions"`
	StatusCounts PermissionStatusCounts `json:"permission_status_counts,omitempty"`
	Principal    Principal              `json:"principal"`
	Start        string                 `json:"start"`
	TotalPages   int                    `json:"total_pages"`
	Warnings     map[string][]string    `json:"warnings"`
}

type Permissions struct {
	Action       string `json:"action,omitempty"`
	Category     string `json:"category"`
	Count        int    `json:"count,omitempty"`
	LastExecuted string `json:"last_executed_date,omitempty"`
	Permission   string `json:"permission"`
	Status       string `json:"status"`
}

type PermissionStatusCounts struct {
	Total      int `json:"total_permission_count"`
	Unassessed int `json:"unassessed_permission_count"`
	Unused     int `json:"unused_permission_count"`
	Used       int `json:"used_permission_count"`
}

type RemediationPolicy struct {
	Policy Policy `json:"policy"`
}

type Policy struct {
	Statement []PolicyDetails `json:"Statement"`
	Version   string          `json:"Version"`
}

type PolicyDetails struct {
	Sid       string   `json:"Sid"`
	NotAction []string `json:"NotAction"`
	Effect    string   `json:"Effect"`
	Resource  []string `json:"Resource"`
}
