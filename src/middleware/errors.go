package middleware

import (
	"errors"
	"scrapher/src/global"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Just patching things up. This'll be over in a jiffy."
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}
	if mongo.IsDuplicateKeyError(err) {
		code = fiber.StatusBadRequest
		message = "Resource already exists"
	}
	return ctx.Status(code).JSON(global.Response[*interface{}]{
		Message: message,
	})
}
