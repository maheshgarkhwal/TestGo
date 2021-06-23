package inter

import (
	"test/database"
	"test/model"
)

type ServiceBook struct {
}

func NewServiceBook() IBook {
	return &ServiceBook{}
}

type IBook interface {
	NewBookService(b model.Book) (model.Book, error)
	UpdateBookService(b model.Book, id string) (model.Book, error)
	GetBookService() ([]model.Book, error)
	DeleteBookService(id string) model.Book
	GetBookByIdService(id string) model.Book
}

func (S *ServiceBook) NewBookService(b model.Book) (model.Book, error) {
	db := database.DBConn

	if err := db.Create(&b).Error; err != nil {
		return b, err
	}
	return b, nil
}

func (S *ServiceBook) UpdateBookService(b model.Book, id string) (model.Book, error) {
	db := database.DBConn
	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		return book, nil
	} else {
		if err := db.Model(&book).Updates(model.Book{Title: b.Title, Rating: b.Rating, Author: b.Author}).Error; err != nil {
			return book, err
		}
		return book, nil
	}
}

func (S *ServiceBook) GetBookService() ([]model.Book, error) {
	db := database.DBConn
	var books []model.Book
	if err := db.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
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
