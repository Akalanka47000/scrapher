package rodext

import (
	"net/http"
	"scrapher/src/global"

	"github.com/gofiber/fiber/v2"
)

type RodErrorDetail struct {
	TargetStatus int    `json:"target_status,omitempty"`
	TargetDetail string `json:"target_detail"`
}

const (
	ErrMsgFailedToAnalyzeWebpage = "Failed to analyze webpage"
)

var (
	ErrTargetIsNotValidHTML = global.NewExtendedFiberError(
		fiber.NewError(http.StatusUnprocessableEntity, ErrMsgFailedToAnalyzeWebpage),
		RodErrorDetail{
			TargetDetail: "The target url is not a valid HTML page",
		},
	)
	ErrConnectionError = global.NewExtendedFiberError(
		fiber.NewError(http.StatusUnprocessableEntity, ErrMsgFailedToAnalyzeWebpage),
		RodErrorDetail{
			TargetDetail: "Connection error, which most likely means that the target url is invalid",
		},
	)
)
