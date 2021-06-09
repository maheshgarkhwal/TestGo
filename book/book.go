package book

import (
	"fmt"
	"test/database"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	fmt.Print(books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.Status(200).JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)

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

	bookData := new(Book)
	c.BodyParser(bookData)
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	} else {
		db.Model(&book).Updates(Book{Title: bookData.Title, Rating: bookData.Rating, Author: bookData.Author})
		c.Status(200).JSON(book)
	}
}

func DeleteBook(c *fiber.Ctx) {
	ID := c.Params("ID")
	db := database.DBConn

	var book Book
	db.First(&book, ID)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
