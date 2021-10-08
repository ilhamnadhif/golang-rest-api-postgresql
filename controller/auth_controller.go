package controller

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-postgresql/helper"
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/service"
	"net/http"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authControllerImpl{
		AuthService: authService,
	}
}

func (controller *authControllerImpl) Login(c *gin.Context) {
	var userRequest web.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		panic(helper.IfValidationError(err))
	}
	user := controller.AuthService.Login(userRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	})
}

func (controller *authControllerImpl) Register(c *gin.Context) {
	var userRequest web.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		panic(helper.IfValidationError(err))
	}
	user := controller.AuthService.Register(userRequest)
	c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   user,
	})
}
