package dto

type PerformAnalysisRequest struct {
	TargetURL string `json:"target_url" validate:"required,url" messages:"Please provide a valid target URL"`
}

type PerformAnalysisResult struct {
	HTMLVersion   string `json:"html_version"`
	PageTitle     string `json:"page_title"`
	HeadingCounts struct {
		H1 int `json:"h1"`
		H2 int `json:"h2"`
		H3 int `json:"h3"`
		H4 int `json:"h4"`
		H5 int `json:"h5"`
		H6 int `json:"h6"`
	} `json:"heading_counts"`
	InternalLinkCount     int  `json:"internal_link_count"`
	ExternalLinkCount     int  `json:"external_link_count"`
	InaccessibleLinkCount int  `json:"inaccessible_link_count"`
	ContainsLoginForm     bool `json:"contains_login_form"`
}
