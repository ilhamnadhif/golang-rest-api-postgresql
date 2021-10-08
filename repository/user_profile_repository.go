package repository

import (
	"golang-rest-api-postgresql/model/domain"
	"gorm.io/gorm"
)

type UserProfileRepository interface {
	Create(tx *gorm.DB, user domain.UserProfile) domain.UserProfile
	Delete(tx *gorm.DB, user domain.UserProfile)
	Update(tx *gorm.DB, user domain.UserProfile) domain.UserProfile
	FindById(tx *gorm.DB, userId int) domain.UserProfile
}

type userProfileRepositoryImpl struct {
}

func NewUserProfileRepository() UserProfileRepository {
	return &userProfileRepositoryImpl{}
}

func (repository *userProfileRepositoryImpl) Create(tx *gorm.DB, user domain.UserProfile) domain.UserProfile {
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return user
}

func (repository *userProfileRepositoryImpl) Delete(tx *gorm.DB, user domain.UserProfile) {
	err := tx.Delete(&user).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
}

func (repository *userProfileRepositoryImpl) Update(tx *gorm.DB, user domain.UserProfile) domain.UserProfile{
	err := tx.Save(&user).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return user
}

func (repository *userProfileRepositoryImpl) FindById(tx *gorm.DB, userProfileId int) domain.UserProfile {
	var userProfile domain.UserProfile
	err := tx.First(&userProfile, userProfileId).Error
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	return userProfile
}
