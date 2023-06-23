package insightcloudsecClient

// APIErrorResponse
type APIErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	ErrorType    string `json:"error_type"`
	Traceback    string `json:"traceback"`
}

// Badge
type Badge struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Badges
type Badges struct {
	Badges []Badge `json:"badges"`
}
