package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/imabg/vrnika/pkg/models"
)

// health check
func HealthCheck(c *fiber.Ctx) error {
	currentTime := time.Now()
	SendResponse(c, http.StatusAccepted, currentTime.String())
	return nil
}

func VerifyOTP(c *fiber.Ctx) error {
	var body models.OTPBody
	err := c.BodyParser(&body)
	if err != nil {
		SendResponse(c, http.StatusBadRequest, err.Error())
	}
	// TODO: Retrieve OTP from Redis
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
