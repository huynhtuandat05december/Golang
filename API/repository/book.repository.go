package repository

import (
	"golang_api/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBook() []models.Book
	GetBookByID(bookID uint64) models.Book
	InsertBook(book models.Book) models.Book
	// UpdateBook(book models.Book) models.Book
	DeleteBook(bookID uint64)
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		connection: db,
	}

}

func (db *bookConnection) GetAllBook() []models.Book {
	var books []models.Book
	db.connection.Preload("User").Find(&books)
	return books
}

func (db *bookConnection) GetBookByID(bookID uint64) models.Book {
	var book models.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}

func (db *bookConnection) InsertBook(book models.Book) models.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}

// func (db *bookConnection) UpdateBook(book models.Book) models.Book {
// 	db.connection.Save(&book)
// 	db.connection.Preload("User").Find(&book)
// 	return book
// }

func (db *bookConnection) DeleteBook(bookID uint64) {
	db.connection.Delete(&models.Book{}, bookID)
}
