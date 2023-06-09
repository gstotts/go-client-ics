package insightcloudsecClient

type NewUser struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email"`
	Username     string `json:"username"`
	AccessLevel  string `json:"access_level,omitempty"`
}

type LocalUser struct {
	NewUser
	Password          string `json:"password"`
	TwoFactorRequired bool   `json:"two_factor_required"`
}

type APIUser struct {
	NewUser
	ExpirationDate int64 `json:"expiration_date"`
}

type SAMLUser struct {
	NewUser
	Domain                 string
	AuthenticationServerID int `json:"authentication_server_id"`
}

type SAMLUserCreateRequest struct {
	NewUser
	AuthenticationType string `json:"authentication_type"`
}

// UserCreateRequest
type UserCreateRequest struct {
	NewUser
	ConfirmPassword string `json:"confirm_password"`
}

// User
type User struct {
	Name                           string   `json:"name"`
	EmailAddress                   string   `json:"email_address"`
	Username                       string   `json:"username"`
	Password                       string   `json:"password"`
	AccessLevel                    string   `json:"access_level"`
	TwoFactorRequired              bool     `json:"two_factor_required"`
	ResourceID                     string   `json:"resource_id"`
	UserID                         int      `json:"user_id"`
	OrganizationAdmin              bool     `json:"organization_admin"`
	DomainAdmin                    bool     `json:"domain_admin"`
	DomainViewer                   bool     `json:"domain_viewer"`
	OrganizationName               string   `json:"organization_name"`
	OrganizationID                 int      `json:"organization_id"`
	TwoFactorEnabled               bool     `json:"two_factor_enabled"`
	Auth                           string   `json:"auth"`
	SessionTTL                     int      `json:"session_ttl"`
	SessionExpiration              string   `json:"session_expiration"`
	SessionTimeoutSeconds          int      `json:"session_timeout_seconds"`
	AuthenticationServerID         int      `json:"authentication_server_id"`
	AuthPluginExists               bool     `json:"auth_plugin_exists"`
	NavigationBlacklist            []string `json:"navigation_blacklist"`
	Theme                          string   `json:"theme"`
	RequirePWReset                 bool     `json:"require_pw_reset"`
	CreateDate                     string   `json:"create_date"`
	AWSDefaultExternalID           string   `json:"aws_default_external_id"`
	ApiKeyGenerationAllowed        bool     `json:"api_key_generation_allowed"`
	Rapid7OrgID                    string   `json:"rapid7_org_id"`
	ActiveApiKeyPresent            bool     `json:"active_api_key_present,omitempty"`
	ConsecutiveFailedLoginAttempts int      `json:"consecutive_failed_login_attempts,omitempty"`
	ConsoleAccessDenied            bool     `json:"console_access_denied,omitempty"`
	Groups                         int      `json:"groups,omitempty"`
	LastLoginTime                  string   `json:"last_login_time,omitempty"`
	OwnedResources                 int      `json:"owned_resources,omitempty"`
	ServerName                     string   `json:"server_name,omitempty"`
	Suspended                      bool     `json:"suspended,omitempty"`
}

// UserList
type UserList struct {
	TotalCount int    `json:"total_count"`
	Users      []User `json:"users"`
}

// MFAStatus
type MFAStatus struct {
	Enabled  bool `json:"enabled"`
	Required bool `json:"required"`
}

// OTPSecret
type OTPSecret struct {
	Secret string `json:"otp_secret"`
}
