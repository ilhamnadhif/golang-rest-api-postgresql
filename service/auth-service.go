package service

import (
	"errors"
	"golang-rest-api-postgresql/model/domain"
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/repository"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(request web.UserRequest) web.UserResponse
	Login(request web.UserRequest) web.UserResponse
}

type authServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewAuthService(db *gorm.DB, userRepository repository.UserRepository) AuthService {
	return &authServiceImpl{
		DB:             db,
		UserRepository: userRepository,
	}
}

func (service *authServiceImpl) Register(request web.UserRequest) web.UserResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	convertUser := domain.User{
		Email:    request.Email,
		Password: request.Password,
	}
	user := service.UserRepository.Create(tx, convertUser)

	return web.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (service *authServiceImpl) Login(request web.UserRequest) web.UserResponse {
	tx := service.DB.Begin()
	defer tx.Commit()

	user := service.UserRepository.FindByEmail(tx, request.Email)

	if request.Password != user.Password {
		panic(errors.New("Password yang anda masukkan salah").Error())
	}

	return web.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}
