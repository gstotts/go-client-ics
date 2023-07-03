package insightcloudsecClient

// Badge
type Badge struct {
	Key          string `json:"key"`
	Value        string `json:"value"`
	AutGenerated bool   `json:"auto_generated,omitempty"`
}

// Badges
type Badges struct {
	Badges []Badge `json:"badges"`
}

type BadgedCloud struct {
	ResourceID string `json:"resource_id"`
	Name       string `json:"name"`
}

type BadgeResourceCount struct {
	ResourceCount []BadgeCount `json:"resource_count"`
}
type BadgeCount struct {
	ResourceID string `json:"resource_id"`
	Count      int    `json:"count"`
}
