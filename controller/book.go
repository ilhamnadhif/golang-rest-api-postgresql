package controller

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/service"
	"net/http"
	"strconv"
)

type BookController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}

type bookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &bookControllerImpl{
		BookService: bookService,
	}
}

func (controller *bookControllerImpl) Create(c *gin.Context) {
	var bookRequest web.BookCreateRequest
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		panic(err.Error())
		return
	}
	book := controller.BookService.Create(bookRequest)
	c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   book,
	})
}

func (controller *bookControllerImpl) Update(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		panic(errorParam.Error())
	}
	var bookRequest web.BookUpdateRequest
	bookRequest.ID = uint(id)

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		panic(err.Error())
	}

	book := controller.BookService.Update(bookRequest)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   book,
	})
}

func (controller *bookControllerImpl) Delete(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		panic(errorParam.Error())
	}

	controller.BookService.Delete(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (controller *bookControllerImpl) FindById(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		panic(errorParam.Error())
	}
	book := controller.BookService.FindById(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   book,
	})
}

func (controller *bookControllerImpl) FindAll(c *gin.Context) {
	books := controller.BookService.FindAll()
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   books,
	})
}
