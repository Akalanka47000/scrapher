package dto

type PerformAnalysisRequest struct {
	TargetURL string `validate:"required,url"`
}

type PerformAnalysisResponse struct{}
