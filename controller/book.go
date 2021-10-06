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
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
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
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errorParam.Error(),
		})
		return
	}
	var bookRequest web.BookUpdateRequest
	bookRequest.ID = uint(id)

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	book, errorService := controller.BookService.Update(bookRequest)
	if errorService != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errorService.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   book,
	})
}

func (controller *bookControllerImpl) Delete(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errorParam.Error(),
		})
		return
	}

	errorService := controller.BookService.Delete(id)
	if errorService != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errorService.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (controller *bookControllerImpl) FindById(c *gin.Context) {
	id, errorParam := strconv.Atoi(c.Param("id"))
	if errorParam != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errorParam.Error(),
		})
		return
	}
	book, errorService := controller.BookService.FindById(id)
	if errorService != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errorService.Error(),
		})
		return
	}
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
