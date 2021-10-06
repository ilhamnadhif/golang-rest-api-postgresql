package repository

import (
	"golang-rest-api-postgresql/model/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(tx *gorm.DB, book domain.Book) domain.Book
	Update(tx *gorm.DB, book domain.Book) domain.Book
	Delete(tx *gorm.DB, book domain.Book)
	FindById(tx *gorm.DB, bookId int) (domain.Book, error)
	FindAll(tx *gorm.DB) []domain.Book
}

type bookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &bookRepositoryImpl{}
}

func (repository *bookRepositoryImpl) Create(tx *gorm.DB, book domain.Book) domain.Book {
	err := tx.Create(&book).Error
	if err != nil {
		tx.Rollback()
	}
	return book
}

func (repository *bookRepositoryImpl) Update(tx *gorm.DB, book domain.Book) domain.Book {
	err := tx.Save(&book).Error
	if err != nil {
		tx.Rollback()
	}
	return book
}

func (repository *bookRepositoryImpl) Delete(tx *gorm.DB, book domain.Book) {
	err := tx.Delete(&book).Error
	if err != nil {
		tx.Rollback()
	}
}

func (repository *bookRepositoryImpl) FindById(tx *gorm.DB, bookId int) (domain.Book, error) {
	var book domain.Book
	err := tx.First(&book, bookId).Error
	if err != nil {
		tx.Rollback()
	}
	return book, err
}

func (repository *bookRepositoryImpl) FindAll(tx *gorm.DB) []domain.Book {
	var books []domain.Book
	err := tx.Find(&books).Error
	if err != nil {
		tx.Rollback()
	}
	return books
}
