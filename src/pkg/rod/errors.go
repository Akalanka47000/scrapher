package rodext

type RodErrorDetail struct {
	TargetStatus int    `json:"target_status,omitempty"`
	TargetDetail string `json:"target_detail"`
}

const (
	ErrFailedToAnalyzeWebpage = "Failed to analyze webpage"
)

const (
	ErrDetailTargetUrlIsNotValidHTML = "The target url is not a valid HTML page"
	ErrDetailConnectionError         = "Connection error, which most likely means that the target url is invalid"
)
