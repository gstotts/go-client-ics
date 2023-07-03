package insightcloudsecClient

// APIErrorResponse
type APIErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	ErrorType    string `json:"error_type"`
	Traceback    string `json:"traceback"`
}
