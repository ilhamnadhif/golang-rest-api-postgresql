package controller

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-postgresql/helper"
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/service"
	"net/http"
	"strconv"
)

type UserProfileController interface {
	FindByID(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type userProfileControllerImpl struct {
	UserProfileService service.UserProfileService
}

func NweUserProfileController(userProfileService service.UserProfileService) UserProfileController {
	return &userProfileControllerImpl{
		UserProfileService: userProfileService,
	}
}

func (controller *userProfileControllerImpl) FindByID(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		panic(errorParam.Error())
	}
	profile := controller.UserProfileService.FindByID(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   profile,
	})
}

func (controller *userProfileControllerImpl) Delete(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		panic(errorParam.Error())
	}
	controller.UserProfileService.Delete(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (controller *userProfileControllerImpl) Create(c *gin.Context) {
	var userProfileCreateRequest web.UserProfileCreateRequest
	err := c.ShouldBindJSON(&userProfileCreateRequest)
	if err != nil {
		panic(helper.IfValidationError(err))
	}
	profile := controller.UserProfileService.Create(userProfileCreateRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   profile,
	})
}

func (controller *userProfileControllerImpl) Update(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		panic(errorParam.Error())
	}
	var userProfileUpdateRequest web.UserProfileUpdateRequest
	userProfileUpdateRequest.ID = id
	err := c.ShouldBindJSON(&userProfileUpdateRequest)
	if err != nil {
		panic(helper.IfValidationError(err))
	}
	profile := controller.UserProfileService.Update(userProfileUpdateRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   profile,
	})
}
