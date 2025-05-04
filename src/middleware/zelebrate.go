package middleware

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

var validate = validator.New()

type ZelebrateSegment string

const (
	ZelebrateSegmentBody    ZelebrateSegment = "body"
	ZelebrateSegmentParams  ZelebrateSegment = "params"
	ZelebrateSegmentQuery   ZelebrateSegment = "query"
	ZelebrateSegmentHeaders ZelebrateSegment = "headers"
)

// Zelebrate is a middleware function that validates one of the body, params, or query of the request
// against the given struct type T.
func Zelebrate[T any](segment ZelebrateSegment) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		target := new(T)
		switch segment {
		case ZelebrateSegmentBody:
			ctx.BodyParser(target)
		case ZelebrateSegmentParams:
			ctx.ParamsParser(target)
		case ZelebrateSegmentQuery:
			ctx.QueryParser(target)
		case ZelebrateSegmentHeaders:
			target = lo.ToPtr(lo.CastJSON[T](ctx.GetReqHeaders()))
		}
		firstFormattedErr := ""
		errs := validate.Struct(target)
		if errs != nil {
			reflectedTarget := reflect.TypeOf(target).Elem()
			for _, err := range lo.Cast[validator.ValidationErrors](errs) {
				field, ok := reflectedTarget.FieldByName(err.StructField())
				if ok {
					messages := field.Tag.Get("messages")
					for message := range strings.SplitSeq(messages, ",") {
						messageSlice := strings.Split(message, "=")
						if len(messageSlice) == 2 && messageSlice[0] == err.Tag() {
							firstFormattedErr = messageSlice[1]
							break
						}
					}
					if firstFormattedErr == "" && messages != "" {
						firstFormattedErr = messages
					}
				}
				if firstFormattedErr == "" {
					firstFormattedErr = fmt.Sprintf("%s failed on the '%s' tag against value '%s'",
						err.Field(), err.Tag(), err.Value())
				}
			}
		}
		if firstFormattedErr != "" {
			panic(fiber.NewError(fiber.StatusUnprocessableEntity, firstFormattedErr))
		}
		return ctx.Next()
	}
}
