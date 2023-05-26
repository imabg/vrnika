package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/imabg/vrnika/cmd/vrnika/database"
	"github.com/imabg/vrnika/pkg/models"
)

func CreateCompany(c *fiber.Ctx) error {
	var company models.Company
	err := c.BodyParser(&company)
	if err != nil {
		SendResponse(c, http.StatusNotAcceptable, err.Error())
	}
	err = company.Create(database.DB)
	if err != nil {
		SendResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Company has created",
		"data":    &company,
	})

}
