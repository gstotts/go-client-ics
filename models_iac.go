package insightcloudsecClient

type IACConfig struct {
	ID                     int      `json:"id"`
	OrganizationID         int      `json:"organization_id"`
	Name                   string   `json:"name"`
	Description            string   `json:"description"`
	PackID                 int      `json:"pack_id"`
	Source                 string   `json:"source"`
	InsightsBlacklist      []string `json:"insights_blacklist"`
	InsightsWarnOnly       []string `json:"insights_warn_only"`
	SlackChannel           string   `json:"slack_channel"`
	EmailRecipients        []string `json:"email_recipients"`
	LastBuild              string   `json:"last_build_at"`
	Created                string   `json:"created_at"`
	Modified               string   `json:"last_modified"`
	TotalBuilds            int      `json:"total_builds"`
	SuccessCount           int      `json:"success_count"`
	FailureCount           int      `json:"failure_count"`
	ConsumptionURL         string   `json:"consumption_url"`
	Favorite               bool     `json:"favorite"`
	TFCCount               int      `json:"tfc_count"`
	LocalExceptionsEnabled bool     `json:"local_exceptions_enabled"`
}
