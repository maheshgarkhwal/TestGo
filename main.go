package main

import (
	"log"
	"test/database"
	route "test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	app := fiber.New()
	route.SetupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}
