package route

import (
	inter "test/interface"

	"github.com/gofiber/fiber/v2"
)

var Service inter.IBook

func SetupRoutes(app *fiber.App, s inter.IBook) {
	Service = s
	//book route
	app.Get("/api/v1/book", GetBooks)
	app.Get("/api/v1/book/:id", GetBookById)
	app.Post("/api/v1/book", NewBook)
	app.Put("/api/v1/book/:id", UpdateBook)
	app.Delete("/api/v1/book/:id", DeleteBook)

	//date insert through excel sheet
	app.Post("/api/v1/xl", DataInsert)

	//sending mail //service.Authentication
	app.Post("/api/v1/mail", Mailer)

	//data transfer through channels
	app.Get("api/v1/channels", Channel)

	//user registration
	app.Post("api/v1/register", Registeration)

	//user login
	app.Post("api/v1/login", Login)

	//GetUser
	app.Get("api/v1/user", GetUser)

	//implementing api using interface
	app.Post("api/v1/ibook", PostStudent)

}
