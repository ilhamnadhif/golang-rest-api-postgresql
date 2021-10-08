package repository

import (
	"golang-rest-api-postgresql/model/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(tx *gorm.DB, book domain.Book) domain.Book
	Update(tx *gorm.DB, book domain.Book) domain.Book
	Delete(tx *gorm.DB, book domain.Book)
	FindById(tx *gorm.DB, bookId int) domain.Book
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
		panic(err.Error())
	}
	return book
}

func (repository *bookRepositoryImpl) Update(tx *gorm.DB, book domain.Book) domain.Book {
	err := tx.Save(&book).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return book
}

func (repository *bookRepositoryImpl) Delete(tx *gorm.DB, book domain.Book) {
	err := tx.Delete(&book).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
}

func (repository *bookRepositoryImpl) FindById(tx *gorm.DB, bookId int) domain.Book {
	var book domain.Book
	err := tx.First(&book, bookId).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return book
}

func (repository *bookRepositoryImpl) FindAll(tx *gorm.DB) []domain.Book {
	var books []domain.Book
	err := tx.Find(&books).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return books
}
