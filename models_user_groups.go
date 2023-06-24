package insightcloudsecClient

type Group struct {
	ID                     int    `json:"group_id"`
	ResourceID             string `json:"resource_id"`
	Name                   string `json:"name"`
	Users                  int    `json:"users"`
	Roles                  int    `json:"roles"`
	EntitlementsConfigured bool   `json:"entitlements_configured"`
}

type Groups struct {
	Groups []Group `json:"groups"`
}

type createGroupRequest struct {
	Name string `json:"group_name"`
}

type groupResponse struct {
	Group Group `json:"group"`
}

type addUsersToGroupRequest struct {
	UserResourceIDs []string `json:"user_resource_ids"`
}

type deleteUserFromGroupRequest struct {
	UserResourceID string `json:"user_resource_id"`
}

type updateGroupRolesRequest struct {
	ResourceIDs []string `json:"resource_ids"`
}

type Entitlements struct {
	Entitlements []Entitlement `json:"entitlements"`
}

type Entitlement struct {
	Namespace string `json:"namespace"`
	Role      string `json:"role"`
}

type UserEntitlement struct {
	Entitlement string `json:"entitlement"`
}

type setEntitlementRequest struct {
	GroupIDs     []int         `json:"group_ids"`
	Entitelments []Entitlement `json:"entitlements"`
}
