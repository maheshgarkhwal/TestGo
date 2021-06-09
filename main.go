package main

import (
	"fmt"
	"test/book"
	"test/database"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("hello world")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book/:id", book.Update)
	app.Delete("/api/v1/book/:id", book.DeleteBook)

}

func initDatabase() {
	var err error
	dsn := "mahesh:Mahesh@g7@tcp(localhost:3306)/crud?parseTime=true"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connection Opened to Database")
	}
	database.DBConn.AutoMigrate(&book.Book{})
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)
}
