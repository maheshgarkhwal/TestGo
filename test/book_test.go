package test

import (
	"test/database"
	"test/model"
	"test/service"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnection(t *testing.T) {
	database.InitDatabase()
}
func TestNewBook(t *testing.T) {
	book := new(model.Book)
	book.Author = "arjunartistic"
	book.Title = "school memories"
	book.Rating = 3
	result := service.NewBookService(book)
	require.NotEmpty(t, result)

	require.Equal(t, book.Author, result.Author)
	require.Equal(t, book.Title, result.Title)
	require.Equal(t, book.Rating, result.Rating)
}

func TestGetBook(t *testing.T) {
	result := service.GetBooksService()
	require.NotEmpty(t, result)
}

func GetBookByIdService(t *testing.T) {
	id := "2"
	result := service.GetBookByIdService(id)
	require.NotEmpty(t, result)
}

func TestDeleteBook(t *testing.T) {
	id := "21"
	result := service.DeleteBookService(id)
	require.NotEmpty(t, result)
}
