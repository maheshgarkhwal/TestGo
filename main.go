package main

import (
	"log"
	"test/database"
	inter "test/interface"
	route "test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	app := fiber.New()
	service := inter.NewServiceBook()
	route.SetupRoutes(app, service)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}
