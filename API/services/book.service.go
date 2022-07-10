package services

import (
	"golang_api/dto"
	"golang_api/models"
	"golang_api/repository"
)

//BookService is a ....
type BookService interface {
	GetAllBook() []models.Book
	GetBookByID(bookID uint64) models.Book
	InsertBook(book dto.BookCreateDTO) models.Book
	// UpdateBook(book dto.BookUpdateDTO) models.Book
	DeleteBook(bookID uint64)
	IsAllowedToEdit(userID uint64, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

//NewBookService .....
func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (service *bookService) GetAllBook() []models.Book {
	res := service.bookRepository.GetAllBook()
	return res
}

func (service *bookService) GetBookByID(bookID uint64) models.Book {
	res := service.bookRepository.GetBookByID(bookID)
	return res
}

func (service *bookService) InsertBook(bookDTO dto.BookCreateDTO) models.Book {
	newBook := models.Book{
		Title:       bookDTO.Title,
		Description: bookDTO.Description,
		UserID:      bookDTO.UserID,
	}
	res := service.bookRepository.InsertBook(newBook)
	return res
}

// func (service *bookService) UpdateBook(bookUpdateDTO dto.BookUpdateDTO) models.Book {

// }

func (service *bookService) DeleteBook(bookID uint64) {
	service.bookRepository.GetBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID uint64, bookID uint64) bool {
	b := service.bookRepository.GetBookByID(bookID)
	return userID == b.UserID
}
