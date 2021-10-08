package repository

import (
	"errors"
	"golang-rest-api-postgresql/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(tx *gorm.DB, user domain.User) domain.User
	Delete(tx *gorm.DB, user domain.User)
	FindByID(tx *gorm.DB, userId int) domain.User
	FindByEmail(tx *gorm.DB, email string) domain.User
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) Create(tx *gorm.DB, user domain.User) domain.User {
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return user
}

func (repository *userRepositoryImpl) Delete(tx *gorm.DB, user domain.User) {
	err := tx.Delete(&user).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
}

func (repository *userRepositoryImpl) FindByID(tx *gorm.DB, userId int) domain.User {
	var user domain.User
	result := tx.Joins("UserProfile").Where("users.id = ?", userId).Find(&user)
	if result.RowsAffected < 1 {
		tx.Rollback()
		panic(errors.New("row tidak ditemukan").Error())
	}
	if result.Error != nil {
		tx.Rollback()
		panic(result.Error.Error())
	}
	return user
}

func (repository *userRepositoryImpl) FindByEmail(tx *gorm.DB, email string) domain.User {
	var user domain.User
	err := tx.Where("email = ?", email).First(&user).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return user
}
