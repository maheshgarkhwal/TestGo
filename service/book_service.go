package service

import (
	"fmt"
	"test/database"
	"test/model"
)

func GetBooksService() []model.Book {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	return books
}

func NewBookService(book *model.Book) *model.Book {
	db := database.DBConn
	db.Create(&book)
	fmt.Print(book)
	return book
}

func UpdateService(id string, bookData *model.Book) model.Book {
	db := database.DBConn
	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		return book
	} else {
		db.Model(&book).Updates(model.Book{Title: bookData.Title, Rating: bookData.Rating, Author: bookData.Author})
		return book
	}
}

func DeleteBookService(id string) model.Book {
	db := database.DBConn
	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		return book
	}
	db.Delete(&book)
	return book
}

func GetBookByIdService(id string) model.Book {
	db := database.DBConn
	var book model.Book
	db.Find(&book, id)
	return book
}
