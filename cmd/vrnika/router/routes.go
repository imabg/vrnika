package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imabg/vrnika/cmd/vrnika/handlers"
)

func SetupRoutes(engine *fiber.App) {
	api := engine.Group("/api")
	// routes
	api.Get("/healthcheck", handlers.HealthCheck)

	// Company routes
	capi := api.Group("/c")

	capi.Post("/create", handlers.CreateCompany)
}
