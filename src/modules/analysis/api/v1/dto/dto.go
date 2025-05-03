package dto

type PerformAnalysisRequest struct {
	TargetURL string `json:"target_url" validate:"required,url" messages:"Please provide a valid target URL"`
}

type PerformAnalysisResponse struct{}
