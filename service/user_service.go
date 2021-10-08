package service

import (
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/repository"
	"gorm.io/gorm"
)

type UserService interface {
	Delete(userId int)
	FindByID(userId int) web.UserWithProfileResponse
}

type userServiceImpl struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
}

func NewUserService(db *gorm.DB, userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		DB:             db,
		UserRepository: userRepository,
	}
}

func (service *userServiceImpl) Delete(userId int) {
	tx := service.DB.Begin()
	defer tx.Commit()

	findUser := service.UserRepository.FindByID(tx, userId)

	service.UserRepository.Delete(tx, findUser)
}

func (service *userServiceImpl) FindByID(userId int) web.UserWithProfileResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	user := service.UserRepository.FindByID(tx, userId)

	var books []web.BookResponse
	for _, book := range user.Books {
		books = append(books, web.BookResponse{
			ID:          book.ID,
			UserID:      book.UserID,
			Title:       book.Title,
			Description: book.Description,
			Price:       book.Price,
			Rating:      book.Rating,
		})
	}

	return web.UserWithProfileResponse{
		ID:    user.ID,
		Email: user.Email,
		UserProfile: web.UserProfileResponse{
			ID:        user.UserProfile.ID,
			UserID:    user.UserProfile.UserID,
			FirstName: user.UserProfile.FirstName,
			LastName:  user.UserProfile.LastName,
			Address:   user.UserProfile.Address,
			Age:       user.UserProfile.Age,
		},
		Books: books,
	}
}
