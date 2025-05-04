package middleware

import (
	"scrapher/src/global"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// List of paths which will be ignored by the logger.
var ZapWhitelists = []string{
	"/system/health",
	"/system/liveness",
	"/system/readiness",
	"/system/metrics",
	"/.well-known/appspecific/com.chrome.devtools.json",
}

func getZapLogFields(c *fiber.Ctx) []zap.Field {
	return []zap.Field{
		zap.String(global.CtxCorrelationID, lo.Cast[string](c.Locals(global.CtxCorrelationID))),
	}
}

func getLogFields(c *fiber.Ctx) []any {
	headers := c.GetReqHeaders()
	return []any{
		"ip", c.IP(),
		"status", c.Response().StatusCode(),
		"method", c.Method(),
		"path", c.Path(),
		"user-agent", lo.FirstOrEmpty(headers[global.HdrUserAgent]),
		"payload", string(c.Body()),
		"query", c.Queries(),
	}
}

// Zapped is a middleware that overrides the default logger with zapcore and sets up an http request logger.
// The logs are populated with the correlation ID associated with the request.
func Zapped(c *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	log.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		ZapOptions: []zap.Option{
			zap.Fields(getZapLogFields(c)...),
		},
	}))

	log.Infow("Request initiated", getLogFields(c)...)

	return fiberzap.New(fiberzap.Config{
		Logger: logger,
		FieldsFunc: func(c *fiber.Ctx) []zap.Field {
			return append(
				getZapLogFields(c),
				zap.String("path", c.Path()),
				zap.Any("query", c.Queries()),
			)
		},
		SkipURIs: ZapWhitelists,
		Fields:   []string{"latency", "status", "ip", "method"},
		Messages: []string{"Server error", "Client error", "Request completed"},
	})(c)
}

// fiberzapPostRecoveryLog is a temporary solution invoked at the default `errorHandler`
// to log the status of failed http requests since the panic + recover flow we use
// doesn't trigger the fiberzap logger on request completion.
func fiberzapPostRecoveryLog(c *fiber.Ctx) {
	log.Errorw("Request failed", getLogFields(c)...)
}
