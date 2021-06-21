package inter

import (
	"fmt"
	"test/database"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

type IBook interface {
	NewBookService() Book
	UpdateBookService(id string) Book
	GetBookService() []Book
	DeleteBookService(id string) Book
}

func (bk Book) NewBookService() Book {
	db := database.DBConn
	db.Create(&bk)
	fmt.Print(bk)
	return bk
}

func (bk Book) UpdateBookService(id string) Book {
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return book
	} else {
		db.Model(&book).Updates(Book{Title: bk.Title, Rating: bk.Rating, Author: bk.Author})
		return book
	}
}

func (bk Book) GetBookService() []Book {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return books
}

func (bk Book) DeleteBookService(id string) Book {
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return book
	}
	db.Delete(&book)
	return book
}
