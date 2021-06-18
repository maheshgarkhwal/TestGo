package main

import (
	"log"
	route "test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	route.SetupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}
