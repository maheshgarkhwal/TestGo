package route

import (
	"test/service"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {

	//book route
	app.Get("/api/v1/book", service.Authentication, service.GetBooks)
	app.Get("/api/v1/book/:id", service.Authentication, service.GetBook)
	app.Post("/api/v1/book", service.Authentication, service.NewBook)
	app.Put("/api/v1/book/:id", service.Authentication, service.Update)
	app.Delete("/api/v1/book/:id", service.Authentication, service.DeleteBook)

	//date insert through excel sheet
	app.Post("/api/v1/xl", service.Authentication, service.DataInsert)

	//sending mail
	app.Post("/api/v1/mail", service.Authentication, service.Mailer)

	//data transfer through channels
	app.Get("api/v1/channels", service.Authentication, service.Channel)

	//user registration
	app.Post("api/v1/register", service.Authentication, service.Registeration)

	//user login
	app.Post("api/v1/login", service.Login)

	//user file upload
}
