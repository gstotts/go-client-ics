package insightcloudsecClient

import (
	"strings"
	"time"
)

type Insight struct {
	All                    int                    `json:"all,omitempty"`
	Author                 string                 `json:"author,omitempty"`
	Bots                   []string               `json:"bots,omitempty"`
	ByApp                  map[string]int         `json:"by_app,omitempty"`
	ByCloud                map[int]map[string]int `json:"by_cloud,omitempty"`
	ByResourceGroup        map[string]int         `json:"by_resource_group,omitempty"`
	ByType                 map[string]int         `json:"by_type,omitempty"`
	CacheUpdatedAt         CacheTime              `json:"cache_updated_at,omitempty"`
	Counts                 InsightCounts          `json:"counts"`
	Disabled               bool                   `json:"disabled,omitempty"`
	Description            string                 `json:"description"`
	Duration               float64                `json:"duration,omitempty"`
	CustomSeverity         int                    `json:"custom_severity"`
	Exemptions             int                    `json:"exemptions,omitempty"`
	Favorited              bool                   `json:"favorited"`
	Filters                []Filter               `json:"filters"`
	ID                     int                    `json:"insight_id"`
	InsertedAt             time.Time              `json:"inserted_at"`
	Metadata               string                 `json:"meta_data"`
	Membership             []string               `json:"membership,omitempty"`
	Name                   string                 `json:"name"`
	Notes                  string                 `json:"notes"`
	Released               string                 `json:"released"`
	ResourceGroupBlacklist []string               `json:"resource_group_blacklist"`
	ResourceTypes          []string               `json:"resource_types"`
	Results                int                    `json:"results,omitempty"`
	RiskLayers             []string               `json:"risk_layers,omitempty"`
	Severity               int                    `json:"severity"`
	Source                 string                 `json:"source"`
	SupportedClouds        []string               `json:"supported_clouds"`
	SupportedIACClouds     []string               `json:"supported_iac_clouds"`
	Tags                   []string               `json:"tags"`
	Total                  int                    `json:"total,omitempty"`
	UpdatedAt              time.Time              `json:"updated_at"`
}

type Filter struct {
	Name   string                 `json:"name"`
	Config map[string]interface{} `json:"config"`
}

type InsightCounts struct {
	All             int                    `json:"all"`
	ByCloud         map[int]map[string]int `json:"by_cloud,omitempty"`
	ByResourceGroup map[string]int         `json:"by_resource_group"`
	ByType          map[string]int         `json:"by_type"`
	CacheUpdatedAt  CacheTime              `json:"cache_updated_at"`
	Duration        string                 `json:"duration"`
	Exemptions      int                    `json:"exemptions"`
	Results         int                    `json:"results"`
	Total           int                    `json:"total"`
}

type CacheTime time.Time

func (c *CacheTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02T03:04:05.000000", value)
	if err != nil {
		return err
	}

	*c = CacheTime(t)
	return nil
}
