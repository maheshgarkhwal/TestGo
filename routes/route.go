package route

import (
	"test/service"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {

	//book route
	app.Get("/api/v1/book", service.GetBooks)
	app.Get("/api/v1/book/:id", service.GetBook)
	app.Post("/api/v1/book", service.NewBook)
	app.Put("/api/v1/book/:id", service.Update)
	app.Delete("/api/v1/book/:id", service.DeleteBook)

	//date insert through excel sheet
	app.Post("/api/v1/xl", service.DataInsert)

}
