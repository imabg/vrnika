package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/imabg/vrnika/cmd/vrnika/database"
	"github.com/imabg/vrnika/cmd/vrnika/router"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load .env %s", err.Error())
	}
}

func main() {
	if err := database.ConnectToDB(); err != nil {
		log.Fatalf("could not connect to DB: %s", err.Error())
	}
	engine := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "vrnika",
		AppName:       "vrnika 0.1",
	})

	router.SetupRoutes(engine)

	engine.Listen(":8080")
}
