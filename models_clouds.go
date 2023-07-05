package insightcloudsecClient

type CloudTypes struct {
	Clouds []CloudType `json:"clouds"`
}

type CloudType struct {
	ID     string `json:"cloud_type_id"`
	Access string `json:"cloud_access"`
	Name   string `json:"name"`
}

type Clouds struct {
	Clouds     []Cloud `json:"clouds"`
	TotalCount int     `json:"total_count"`
}

type Cloud struct {
	client                      *Client
	ID                          int                        `json:"id"`
	Name                        string                     `json:"name"`
	TypeID                      string                     `json:"cloud_type_id"`
	AccountID                   string                     `json:"account_id"`
	CreationTime                string                     `json:"creation_time"`
	HostAssessmentStatus        []string                   `json:"host_assessment_status"`
	HostAssessmentEnabled       bool                       `json:"host_assessment_enabled"`
	Status                      string                     `json:"status"`
	BadgeCount                  int                        `json:"badge_count"`
	ResourceCount               int                        `json:"resource_count"`
	FailedResourceTypes         []CloudResourceTypeFailure `json:"failed_resource_types,omitempty"`
	LastRefereshed              string                     `json:"last_refreshed"`
	RoleArn                     string                     `json:"role_arn"`
	GroupResourceID             string                     `json:"group_resource_id"`
	ResourceID                  string                     `json:"resource_id"`
	EDHRole                     string                     `json:"event_driven_harvest_role"`
	StrategyID                  int                        `json:"strategy_id"`
	CloudOrganizationID         string                     `json:"cloud_organization_id,omitempty"`
	CloudOrganizationDomainName string                     `json:"cloud_organization_domain_name,omitempty"`
	CloudOrganizationNickname   string                     `json:"cloud_organization_nickname,omitempty"`
}

type CloudResourceTypeFailure struct {
	ResourceType string   `json:"resource_type"`
	Permissions  []string `json:"permissions"`
}
