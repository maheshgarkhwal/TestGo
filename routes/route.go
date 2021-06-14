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

	//sending mail
	app.Post("/api/v1/mail", service.Mailer)

	//data transfer through channels
	app.Get("api/v1/channels", service.Channel)

	//user registration
	app.Post("api/v1/register", service.Registeration)

	//user login
	app.Post("api/v1/login", service.Login)

	//token verification
	app.Post("api/v1/auth", service.Authentication)

	//user file upload
}
