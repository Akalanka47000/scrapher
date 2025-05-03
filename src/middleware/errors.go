package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"scrapher/src/global"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Just patching things up. This'll be over in a jiffy."
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}
	return ctx.Status(code).JSON(global.Response[any]{
		Message: message,
	})
}
