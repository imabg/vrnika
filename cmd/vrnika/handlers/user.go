package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/imabg/vrnika/cmd/vrnika/database"
	"github.com/imabg/vrnika/pkg/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		SendResponse(c, http.StatusNotAcceptable, err.Error())
	}
	err = user.Create(database.DB)
	if err != nil {
		SendResponse(c, http.StatusBadRequest, err.Error())
	}
	// Generate and Send OTP
	var otp models.OTP
	err = otp.Save(database.DB)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "User has created",
		"data":    otp.OTP,
	})
}
