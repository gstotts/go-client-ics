package insightcloudsecClient

// UserInfoResponse
type UserInfoResponse struct {
	ResourceID              string   `json:"resource_id"`
	Name                    string   `json:"name"`
	UserID                  int      `json:"user_id"`
	OrganizationAdmin       bool     `json:"organization_admin"`
	DomainAdmin             bool     `json:"domain_admin"`
	DomainViewer            bool     `json:"domain_viewer"`
	EmailAddress            string   `json:"email_address"`
	Username                string   `json:"username"`
	OrganizationName        string   `json:"organization_name"`
	OrganizationID          int      `json:"organization_id"`
	TwoFactorEnabled        bool     `json:"two_factor_enabled"`
	TwoFactorRequired       bool     `json:"two_factor_required"`
	Auth                    string   `json:"auth"`
	SessionTTL              int      `json:"session_ttl"`
	SessionExpiration       string   `json:"session_expiration"`
	SessionTimeoutSeconds   int      `json:"session_timeout_seconds"`
	AuthenticationServerID  int      `json:"authentication_server_id"`
	AuthPluginExists        bool     `json:"auth_plugin_exists"`
	NavigationBlacklist     []string `json:"navigation_blacklist"`
	Theme                   string   `json:"theme"`
	RequirePWReset          bool     `json:"require_pw_reset"`
	CreateDate              string   `json:"create_date"`
	AWSDefaultExternalID    string   `json:"aws_default_external_id"`
	ApiKeyGenerationAllowed bool     `json:"api_key_generation_allowed"`
	Rapid7OrgID             string   `json:"rapid7_org_id"`
}

// UserListResponse
type UserListResponse struct {
	TotalCount int            `json:"total_count"`
	Users      []UserListItem `json:"users"`
}

// UserResponse
type UserListItem struct {
	UserInfoResponse
	ActiveApiKeyPresent            bool   `json:"active_api_key_present"`
	ConsecutiveFailedLoginAttempts int    `json:"consecutive_failed_login_attempts"`
	ConsoleAccessDenied            bool   `json:"console_access_denied"`
	Groups                         int    `json:"groups"`
	LastLoginTime                  string `json:"last_login_time"`
	OwnedResources                 int    `json:"owned_resources"`
	ServerName                     string `json:"server_name"`
	Suspended                      bool   `json:"suspended"`
}
