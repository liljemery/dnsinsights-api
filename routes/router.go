package routes

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"

	docs "github.com/jeremyinoa/dnsinsight-api/docs"
	"github.com/jeremyinoa/dnsinsight-api/internal/middlewares"
	"github.com/jeremyinoa/dnsinsight-api/internal/modules/common"
	"github.com/jeremyinoa/dnsinsight-api/internal/adapters/http/dns"
	"github.com/jeremyinoa/dnsinsight-api/internal/adapters/http/ip"
	"github.com/jeremyinoa/dnsinsight-api/internal/adapters/http/network"
	"github.com/jeremyinoa/dnsinsight-api/configs"
)

// @title DNS Insight API
// @version 0.1.0
// @description Open-source API for DNS and network diagnostics.
// @BasePath /
func Initialize(cfg *configs.Config) *fiber.App {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	// Ensure Swagger uses correct base path
	docs.SwaggerInfo.BasePath = "/"

	// Middlewares
	app.Use(middlewares.Recovery())
	app.Use(middlewares.Logger())

	// Swagger docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// Health/meta
	app.Get("/api/meta", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(common.Response{
			Status:  "success",
			Message: "meta",
			Data: fiber.Map{
				"name":    cfg.AppName,
				"env":     cfg.AppEnv,
				"version": "0.1.0",
			},
		})
	})

	dnsHandler := dns.NewHandler()
	ipHandler := ip.NewHandler()
	netHandler := network.NewHandler()

	api := app.Group("/api")
	public := api.Group("/public", middlewares.RateLimit(cfg.RateLimit))

	// DNS
	public.Get("/dns/resolve", dnsHandler.Resolve)
	public.Get("/ip/reverse", dnsHandler.Reverse)
	public.Get("/network/ttl", dnsHandler.TTL)

	// IP
	public.Get("/ip/lookup", ipHandler.Lookup)

	// Network
	public.Get("/network/ping", netHandler.Ping)

	return app
}
