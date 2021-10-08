package controller

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/service"
	"net/http"
	"strconv"
)

type UserController interface {
	FindByID(c *gin.Context)
	Delete(c *gin.Context)
}

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userControllerImpl{
		UserService: userService,
	}
}

func (controller *userControllerImpl) FindByID(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		panic(errorParam.Error())
	}

	user := controller.UserService.FindByID(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	})
}

func (controller *userControllerImpl) Delete(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		panic(errorParam.Error())
	}

	controller.UserService.Delete(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
