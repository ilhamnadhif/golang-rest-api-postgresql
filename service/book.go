package service

import (
	"golang-rest-api-postgresql/helper"
	"golang-rest-api-postgresql/model/domain"
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/repository"
	"gorm.io/gorm"
)

type BookService interface {
	Create(request web.BookCreateRequest) web.BookResponse
	Update(request web.BookUpdateRequest) (web.BookResponse, error)
	Delete(bookId int) error
	FindById(bookId int) (web.BookResponse, error)
	FindAll() []web.BookResponse
}

type bookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *gorm.DB
}

func NewBookService(db *gorm.DB, bookRepository repository.BookRepository) BookService {
	return &bookServiceImpl{
		DB:             db,
		BookRepository: bookRepository,
	}
}

func (service *bookServiceImpl) Create(request web.BookCreateRequest) web.BookResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	book := domain.Book{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Rating:      request.Rating,
	}
	book = service.BookRepository.Create(tx, book)
	return helper.ToBookResponse(book)
}

func (service *bookServiceImpl) Update(request web.BookUpdateRequest) (web.BookResponse, error) {
	tx := service.DB.Begin()
	defer tx.Commit()

	findBook, err := service.BookRepository.FindById(tx, int(request.ID))
	if err != nil {
		return helper.ToBookResponse(findBook), err
	}

	book := domain.Book{
		ID:          findBook.ID,
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Rating:      request.Rating,
	}
	book = service.BookRepository.Update(tx, book)
	return helper.ToBookResponse(book), nil
}

func (service *bookServiceImpl) Delete(bookId int) error {
	tx := service.DB.Begin()
	defer tx.Commit()

	findBook, err := service.BookRepository.FindById(tx, bookId)
	if err != nil {
		return err
	}

	service.BookRepository.Delete(tx, findBook)
	return nil
}

func (service *bookServiceImpl) FindById(bookId int) (web.BookResponse, error) {
	tx := service.DB.Begin()
	defer tx.Commit()

	book, err := service.BookRepository.FindById(tx, bookId)
	if err != nil {
		return helper.ToBookResponse(book), err
	}
	return helper.ToBookResponse(book), nil
}

func (service *bookServiceImpl) FindAll() []web.BookResponse {
	tx := service.DB.Begin()
	defer tx.Commit()
	books := service.BookRepository.FindAll(tx)
	return helper.ToBookResponses(books)
}
