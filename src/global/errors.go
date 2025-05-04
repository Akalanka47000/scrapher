package global

import "github.com/gofiber/fiber/v2"

type ExtendedFiberError struct {
	BaseError *fiber.Error
	Detail    any `json:"detail"`
}

// Creates a new ExtendedFiberError with the given base error and detail.
func NewExtendedFiberError(baseError *fiber.Error, detail any) *ExtendedFiberError {
	return &ExtendedFiberError{
		BaseError: baseError,
		Detail:    detail,
	}
}

// Error implements the error interface for ExtendedFiberError. Do not rename this method.
func (e *ExtendedFiberError) Error() string {
	return e.BaseError.Error()
}
