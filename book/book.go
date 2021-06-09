package book

import (
	"fmt"
	"test/database"
	"test/model"

	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	fmt.Print(books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.Find(&book, id)
	c.Status(200).JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(book.Title)
	fmt.Println(book.Author)

	fmt.Print(book)
	db.Create(&book)
	c.Status(200).JSON(book)
}

func Update(c *fiber.Ctx) {

	bookData := new(model.Book)
	c.BodyParser(bookData)
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	} else {
		db.Model(&book).Updates(model.Book{Title: bookData.Title, Rating: bookData.Rating, Author: bookData.Author})
		c.Status(200).JSON(book)
	}
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
