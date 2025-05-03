package global

import "github.com/gofiber/fiber/v2"

type Response[T any] struct {
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
	Error   *any   `json:"error,omitempty"`
}

type ExtendedFiberError struct {
	BaseError *fiber.Error
	Detail    any `json:"detail"`
}

// Error implements the error interface for ExtendedFiberError. Do not rename this method.
func (e *ExtendedFiberError) Error() string {
	return e.BaseError.Error()
}

func NewExtendedFiberError(baseError *fiber.Error, detail any) *ExtendedFiberError {
	return &ExtendedFiberError{
		BaseError: baseError,
		Detail:    detail,
	}
}
