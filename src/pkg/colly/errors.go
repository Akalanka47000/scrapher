package collyext

type CollyErrorDetail struct {
	TargetStatus int    `json:"target_status,omitempty"`
	TargetDetail string `json:"target_detail"`
}

const (
	ErrFailedToAnalyzeTargetURL = "Failed to analyze target URL"
)

const (
	ErrDetailConnectionError = "Connection error, which most likely means that the target url is invalid"
)
