package insightcloudsecClient

type AuthenticationServer struct {
	ID           int    `json:"server_id"`
	Name         string `json:"server_name"`
	Host         string `json:"server_host"`
	Port         int    `json:"server_port"`
	Secure       int    `json:"secure"`
	Type         string `json:"server_type"`
	GlobalScope  bool   `json:"global_scope"`
	MappedGroups int    `json:"mapped_groups"`
}

type AuthenticationServers struct {
	Servers []AuthenticationServer `json:"servers"`
}
