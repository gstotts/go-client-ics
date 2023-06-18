package insightcloudsecClient

// LocalUser
type LocalUser struct {
	Name              string `json:"name"`
	EmailAddress      string `json:"email"`
	Username          string `json:"username"`
	AccessLevel       string `json:"access_level,omitempty"`
	TwoFactorRequired bool   `json:"two_factor_required"`
}

// APIUser
type APIUser struct {
	Name           string `json:"name"`
	EmailAddress   string `json:"email"`
	Username       string `json:"username"`
	ExpirationDate int64  `json:"expiration_date,omitempty"`
}

// SAMLUser
type SAMLUser struct {
	Name                   string `json:"name"`
	EmailAddress           string `json:"email"`
	Username               string `json:"username"`
	AccessLevel            string `json:"access_level,omitempty"`
	Domain                 string
	AuthenticationServerID int `json:"authentication_server_id"`
}

// UserCreateRequest
type userCreateRequest struct {
	Name              string `json:"name"`
	EmailAddress      string `json:"email"`
	Username          string `json:"username"`
	AccessLevel       string `json:"access_level,omitempty"`
	TwoFactorRequired bool   `json:"two_factor_required"`
}

// UserCreateResponse
type userTempPasswordResponse struct {
	Name             string `json:"name"`
	UserID           int    `json:"user_id"`
	Username         string `json:"username"`
	OrganizationID   int    `json:"organization_id"`
	TemporaryPW      string `json:"temporary_pw"`
	TempPWExpiration string `json:"temp_pw_expiration"`
}

// UserCreateAPIKeyResponse
type userCreateAPIKeyResponse struct {
	Name           string `json:"name"`
	UserID         int    `json:"user_id"`
	Username       string `json:"username"`
	OrganizationID int    `json:"organization_id"`
	ApiKey         string `json:"api_key"`
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
	TemporaryPW                    string   `json:"temporary_pw,omitempty"`
	TempPWExpiration               string   `json:"temp_pw_expiration,omitempty"`
	ApiKey                         string   `json:"api_key,omitempty"`
}

// UserList
type UserList struct {
	TotalCount int    `json:"total_count,omitempty"`
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
