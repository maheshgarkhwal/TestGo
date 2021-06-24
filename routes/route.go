package route

import (
	inter "test/interface"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var Service inter.IBook
var Valid *validator.Validate

func SetupRoutes(a *fiber.App, s inter.IBook) {
	Service = s

	app := a.Group("/api/v1")
	//book route
	app.Get("/book", GetBooks)
	app.Get("book/:id", GetBookById)
	app.Post("/book", NewBook)
	app.Put("/book/:id", UpdateBook)
	app.Delete("/book/:id", DeleteBook)

	//date insert through excel sheet
	app.Post("/xl", DataInsert)

	//sending mail //service.Authentication
	app.Post("/mail", Mailer)

	//data transfer through channels
	app.Get("/channels", Channel)

	//user registration
	app.Post("/register", Registeration)

	//user login
	app.Post("/login", Login)

	//GetUser
	app.Get("/user", GetUser)

	//implementing api using interface
	app.Post("/ibook", PostStudent)

}
