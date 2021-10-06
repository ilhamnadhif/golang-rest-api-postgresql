package main

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-postgresql/config"
	"golang-rest-api-postgresql/controller"
	"golang-rest-api-postgresql/repository"
	"golang-rest-api-postgresql/service"
)

func main() {
	db := config.SetupDatabaseConnection()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(db, bookRepository)
	bookController := controller.NewBookController(bookService)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	book := v1.Group("/books")
	book.POST("/", bookController.Create)
	book.PUT("/:id", bookController.Update)
	book.DELETE("/:id", bookController.Delete)
	book.GET("/", bookController.FindAll)
	book.GET("/:id", bookController.FindById)

	r.Run(":3000")
}