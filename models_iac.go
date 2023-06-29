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
	Author                string                            `json:"author_name"`
	BuildID               int                               `json:"id"`
	ConfigID              int                               `json:"config_id"`
	ConfigName            string                            `json:"config_name"`
	CreateTime            string                            `json:"create_time"`
	Details               ScanDetails                       `json:"details"`
	Duration              int                               `json:"duration"`
	Errors                []string                          `json:"errors"`
	FailedInsights        int                               `json:"failed_insights"`
	FailedResources       int                               `json:"failed_resources"`
	FailedResourcesByType []ResourcesByType                 `json:"failed_resources_by_type"`
	IACProvider           string                            `json:"iac_provider"`
	Pack                  PackInfo                          `json:"pack"`
	PassedResources       int                               `json:"passed_resources"`
	ResourceMapping       map[string]map[string]interface{} `json:"resource_mapping"`
	ResourcesByType       []ResourcesByType                 `json:"resources_by_type"`
	ScanName              string                            `json:"scan_name"`
	ScanResults           string                            `json:"scan_results"`
	ScanType              string                            `json:"scan_type"`
	Stacktrace            []string                          `json:"stacktrace"`
	Status                string                            `json:"status"`
	StatusMessage         string                            `json:"status_message"`
	Success               bool                              `json:"success"`
	TotalResources        int                               `json:"total_resources"`
	WarnedInsights        int                               `json:"warned_insights"`
	WarnedResources       int                               `json:"warned_resources"`
	WarnedResourcesByType []ResourcesByType                 `json:"warned_resources_by_type"`
	Warnings              []string                          `json:"warnings"`
}

type ResourcesByType struct {
	Compute []string `json:"cat_compute,omitempty"`
	IAM     []string `json:"cat_iam,omitempty"`
	Network []string `json:"cat_network,omitempty"`
}

type PackInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

type ScanDetails struct {
	FailedInsights   []InsightScanResults `json:"failed_insights"`
	PassedInsights   []InsightScanResults `json:"passed_insights"`
	ScanTemplate     string               `json:"scan_template"`
	FailedResources  int                  `json:"failed_resources"`
	PassedResources  int                  `json:"passed_resources"`
	SkippedInsights  []InsightScanResults `json:"skipped_insights"`
	SkippedResources int                  `json:"skipped_resources"`
	TotalInsights    int                  `json:"total_insights"`
	TotalResources   int                  `json:"total_resources"`
	WarnedInsights   []InsightScanResults `json:"warned_insights"`
}

type InsightScanResults struct {
	ID            int                        `json:"id"`
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	IACSupport    string                     `json:"iac_support"`
	Notes         string                     `json:"notes"`
	ResourceTypes []ScanResultsResourceTypes `json:"resource_types"`
	Severity      int                        `json:"severity"`
	Source        string                     `json:"source"`
	Success       int                        `json:"success"`
	Failure       int                        `json:"failure"`
	Warning       int                        `json:"warning"`
	WarnOnly      bool                       `json:"warn_only"`
}

type ScanResultsResourceTypes struct {
	CloudType    string `json:"cloud_type_id"`
	ResourceType string `json:"resource_type"`
	Supported    bool   `json:"supported"`
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
