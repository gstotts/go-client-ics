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

type Scan struct {
	Name     string `json:"scan_name"`
	Template string `json:"scan_template"`
	Config   string `json:"config_name"`
	Author   string `json:"author_name"`
	Provider string `json:"iac_provider"`
}

type ScanResults struct {
	BuildID         int                    `json:"build_id"`
	Config          string                 `json:"config_name"`
	Details         ScanDetails            `json:"details,omitempty"`
	Errors          []string               `json:"errors"`
	Message         string                 `json:"message"`
	Resources       []ScanResource         `json:"resources,omitempty"`
	ResourceMapping map[string]interface{} `json:"resource_mapping,omitempty"`
	ResultsURL      string                 `json:"scan_results"`
	Stacktrace      []string               `json:"stacktrace"`
	Status          string                 `json:"status"`
	Success         bool                   `json:"success"`
}

type ScanDetails struct {
	FailedInsights   []InsightScanResults `json:"failed_insights"`
	FailedResources  int                  `json:"failed_resources"`
	PassedInsights   []InsightScanResults `json:"passed_insights"`
	PassedResources  int                  `json:"passed_resources"`
	SkippedInsights  []InsightScanResults `json:"skipped_insights"`
	SkippedResources int                  `json:"skipped_resources"`
	TotalInsights    int                  `json:"total_insights"`
	TotalResources   int                  `json:"total_resources"`
	WarnedInsights   []InsightScanResults `json:"warned_insights"`
	WarnedResources  int                  `json:"warned_resources"`
}

type InsightScanResults struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Notes         string   `json:"notes"`
	ResourceTypes []string `json:"resource_types"`
	Severity      int      `json:"severity"`
	Source        string   `json:"source"`
	Success       []string `json:"success"`
	Failure       []string `json:"failure"`
	Warning       []string `json:"warning"`
	WarnOnly      bool     `json:"warn_only"`
}

type ScanResource struct {
	Source       string     `json:"resource_source"`
	Address      string     `json:"resource_address"`
	Name         string     `json:"resource_name"`
	Type         string     `json:"resource_type"`
	Region       string     `json:"resource_region"`
	PassedRules  []ScanRule `json:"passed_rules"`
	WarningRules []ScanRule `json:"warning_rules"`
	FailedRules  []ScanRule `json:"failed_rules"`
}

type ScanRule struct {
	Name        string `json:"rule_name"`
	Description string `json:"rule_description"`
	Detail      string `json:"rule_detail"`
}
