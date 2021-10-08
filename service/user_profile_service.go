package service

import (
	"golang-rest-api-postgresql/model/domain"
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/repository"
	"gorm.io/gorm"
)

type UserProfileService interface {
	Delete(userId int)
	FindByID(userId int) web.UserProfileResponse
	Create(request web.UserProfileCreateRequest) web.UserProfileResponse
	Update(request web.UserProfileUpdateRequest) web.UserProfileResponse
}

type userProfileServiceImpl struct {
	DB                    *gorm.DB
	UserProfileRepository repository.UserProfileRepository
}

func NewUserProfileService(db *gorm.DB, userProfileRepository repository.UserProfileRepository) UserProfileService {
	return &userProfileServiceImpl{
		DB:                    db,
		UserProfileRepository: userProfileRepository,
	}
}

func (service *userProfileServiceImpl) Delete(userId int) {
	tx := service.DB.Begin()
	defer tx.Commit()

	findUser := service.UserProfileRepository.FindById(tx, userId)
	service.UserProfileRepository.Delete(tx, findUser)
}

func (service *userProfileServiceImpl) FindByID(userId int) web.UserProfileResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	user := service.UserProfileRepository.FindById(tx, userId)
	return web.UserProfileResponse{
		ID:        user.ID,
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Address:   user.Address,
		Age:       user.Age,
	}
}

func (service *userProfileServiceImpl) Create(request web.UserProfileCreateRequest) web.UserProfileResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	convertUser := domain.UserProfile{
		UserID:    request.UserID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
		Age:       request.Age,
	}

	user := service.UserProfileRepository.Create(tx, convertUser)
	return web.UserProfileResponse{
		ID:        user.ID,
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Address:   user.Address,
		Age:       user.Age,
	}
}

func (service *userProfileServiceImpl) Update(request web.UserProfileUpdateRequest) web.UserProfileResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	findUser := service.UserProfileRepository.FindById(tx, request.ID)

	convertUser := domain.UserProfile{
		UserID:    findUser.UserID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
		Age:       request.Age,
	}

	user := service.UserProfileRepository.Update(tx, convertUser)
	return web.UserProfileResponse{
		ID:        user.ID,
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Address:   user.Address,
		Age:       user.Age,
	}
}
