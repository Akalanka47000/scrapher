package middleware

import (
	"errors"
	"runtime/debug"
	"scrapher/src/global"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Just patching things up. This'll be over in a jiffy."
	var errorDetail *any
	var e *fiber.Error
	var ee *global.ExtendedFiberError
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	} else if errors.As(err, &ee) {
		code = ee.BaseError.Code
		message = ee.BaseError.Message
		errorDetail = &ee.Detail
	}
	err = ctx.Status(code).JSON(global.Response[any]{
		Message: message,
		Error:   errorDetail,
	})

	fiberzapPostRecoveryLog(ctx)

	return err
}

func StackTraceHandler(ctx *fiber.Ctx, err any) {
	if _, ok := err.(*fiber.Error); ok {
		return
	}
	if _, ok := err.(*global.ExtendedFiberError); ok {
		return
	}
	log.Errorw(lo.Cast[error](err).Error(), "stacktrace", string(debug.Stack()))
}
