package helper

import (
	"golang-rest-api-postgresql/model/domain"
	"golang-rest-api-postgresql/model/web"
)

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		ID:          book.ID,
		UserID:      book.UserID,
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		Rating:      book.Rating,
	}
}

func ToBookResponses(books []domain.Book) []web.BookResponse {
	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponse := web.BookResponse{
			ID:          book.ID,
			UserID:      book.UserID,
			Title:       book.Title,
			Description: book.Description,
			Price:       book.Price,
			Rating:      book.Rating,
		}
		bookResponses = append(bookResponses, bookResponse)
	}
	return bookResponses
}
