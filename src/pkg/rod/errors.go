package rodext

type RodErrorDetail struct {
	TargetStatus int    `json:"target_status,omitempty"`
	TargetDetail string `json:"target_detail"`
}

const (
	ErrFailedToAnalyzeTargetURL = "Failed to analyze target URL"
)

const (
	ErrDetailTargetUrlIsNotValidHTML = "The target url is not a valid HTML page"
	ErrDetailConnectionError         = "Connection error, which most likely means that the target url is invalid"
)
