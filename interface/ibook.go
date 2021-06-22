package inter

import (
	"fmt"
	"test/database"
	"test/model"
)

type ServiceBook struct {
}

func NewServiceBook() IBook {
	return &ServiceBook{}
}

type IBook interface {
	NewBookService(b model.Book) model.Book
	UpdateBookService(b model.Book, id string) model.Book
	GetBookService() []model.Book
	DeleteBookService(id string) model.Book
	GetBookByIdService(id string) model.Book
}

func (S *ServiceBook) NewBookService(b model.Book) model.Book {
	db := database.DBConn
	db.Create(&b)
	fmt.Print(b)
	return b
}

func (S *ServiceBook) UpdateBookService(b model.Book, id string) model.Book {
	db := database.DBConn
	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		return book
	} else {
		db.Model(&book).Updates(model.Book{Title: b.Title, Rating: b.Rating, Author: b.Author})
		return book
	}
}

func (S *ServiceBook) GetBookService() []model.Book {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	return books
}

func (S *ServiceBook) DeleteBookService(id string) model.Book {
	db := database.DBConn
	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		return book
	}
	db.Delete(&book)
	return book
}

func (S *ServiceBook) GetBookByIdService(id string) model.Book {
	db := database.DBConn
	var book model.Book
	db.Find(&book, id)
	return book
}
