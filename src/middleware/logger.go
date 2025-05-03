package middleware

import (
	"scrapher/src/global"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func getLogFields(c *fiber.Ctx) []zap.Field {
	return []zap.Field{
		zap.String(global.CtxCorrelationID, lo.Cast[string](c.Locals(global.CtxCorrelationID))),
	}
}

// Zapped is a middleware that overrides the default logger with zapcore and sets up an http request logger.
// The logs are populated with the correlation ID associated with the request.
func Zapped(c *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	log.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		ZapOptions: []zap.Option{
			zap.Fields(getLogFields(c)...),
		},
	}))

	return fiberzap.New(fiberzap.Config{
		Logger: logger,
		FieldsFunc: func(c *fiber.Ctx) []zap.Field {
			headers := c.GetReqHeaders()
			return append(
				getLogFields(c),
				zap.Any("user-agent", lo.FirstOrEmpty(headers[global.HdrUserAgent])),
			)
		},
	})(c)
}
