package middleware

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

func RegisterMetrics(app *fiber.App) {
	prometheus := fiberprometheus.New(app.Config().AppName)

	path := "/system/metrics"

	app.Use(path, Sentinel)

	prometheus.RegisterAt(app, path)

	prometheus.SetSkipPaths(ZapWhitelists)

	app.Use(prometheus.Middleware)
}
