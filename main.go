package main

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-postgresql/config"
	"golang-rest-api-postgresql/controller"
	"golang-rest-api-postgresql/model/web"
	"golang-rest-api-postgresql/repository"
	"golang-rest-api-postgresql/service"
	"net/http"
)

func main() {
	db := config.SetupDatabaseConnection()

	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(db, bookRepository)
	bookController := controller.NewBookController(bookService)

	userRepository := repository.NewUserRepository()

	authService := service.NewAuthService(db, userRepository)
	authController := controller.NewAuthController(authService)

	userService := service.NewUserService(db, userRepository)
	userController := controller.NewUserController(userService)

	userProfileRepository := repository.NewUserProfileRepository()
	userProfileService := service.NewUserProfileService(db, userProfileRepository)
	userProfileController := controller.NweUserProfileController(userProfileService)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err,
		})
	}))

	v1 := r.Group("/api/v1")

	book := v1.Group("/books")
	book.POST("/", bookController.Create)
	book.PUT("/:id", bookController.Update)
	book.DELETE("/:id", bookController.Delete)
	book.GET("/", bookController.FindAll)
	book.GET("/:id", bookController.FindById)

	auth := v1.Group("/auth")
	auth.POST("/login", authController.Login)
	auth.POST("/register", authController.Register)

	user := v1.Group("/user")
	user.GET("/:id", userController.FindByID)
	user.DELETE("/:id", userController.Delete)

	userProfile := v1.Group("/profile")
	userProfile.GET("/:id", userProfileController.FindByID)
	userProfile.POST("/", userProfileController.Create)
	userProfile.PUT("/:id", userProfileController.Update)
	userProfile.DELETE("/:id", userProfileController.Delete)

	r.Run(":4000")
}
