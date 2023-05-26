package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// health check
func HealthCheck(c *fiber.Ctx) error {
	currentTime := time.Now()
	SendResponse(c, http.StatusAccepted, currentTime.String())
	return nil
}

// send response
func SendResponse(c *fiber.Ctx, code int, data interface{}) error {
	c.Status(code)
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}
