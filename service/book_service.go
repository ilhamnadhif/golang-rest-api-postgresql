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
	Update(request web.BookUpdateRequest) web.BookResponse
	Delete(bookId int)
	FindById(bookId int) web.BookResponse
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
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
		Price:       uint(request.Price),
		Rating:      request.Rating,
	}
	book = service.BookRepository.Create(tx, book)
	return helper.ToBookResponse(book)
}

func (service *bookServiceImpl) Update(request web.BookUpdateRequest) web.BookResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	findBook := service.BookRepository.FindById(tx, int(request.ID))

	convertBook := domain.Book{
		ID:          findBook.ID,
		Title:       request.Title,
		Description: request.Description,
		Price:       uint(request.Price),
		Rating:      request.Rating,
	}
	book := service.BookRepository.Update(tx, convertBook)
	return helper.ToBookResponse(book)
}

func (service *bookServiceImpl) Delete(bookId int) {
	tx := service.DB.Begin()
	defer tx.Commit()

	findBook := service.BookRepository.FindById(tx, bookId)

	service.BookRepository.Delete(tx, findBook)
}

func (service *bookServiceImpl) FindById(bookId int) web.BookResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	book := service.BookRepository.FindById(tx, bookId)
	return helper.ToBookResponse(book)
}

func (service *bookServiceImpl) FindAll() []web.BookResponse {
	tx := service.DB.Begin()
	defer tx.Commit()
	books := service.BookRepository.FindAll(tx)
	return helper.ToBookResponses(books)
}
