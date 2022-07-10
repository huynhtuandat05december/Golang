package main

import (
	"golang_api/config"
	"golang_api/controllers"
	"golang_api/middlewares"
	"golang_api/repository"
	"golang_api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
	//repository
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)
	//service
	jwtService  services.JWTService  = services.NewJWTService()
	authService services.AuthService = services.NewAuthService(userRepository)
	userService services.UserService = services.NewUserService(userRepository)
	bookService services.BookService = services.NewBookService(bookRepository)
	//controller
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController = controllers.NewUserController(userService)
	bookController controllers.BookController = controllers.NewBookController(bookService)
)

func main() {
	defer config.CloseConnectionDB(db)
	r := gin.Default()

	authRouters := r.Group("api/v1/auth")
	{
		authRouters.POST("/login", authController.Login)
		authRouters.POST("/register", authController.Register)
	}

	userRouters := r.Group("api/v1/user", middlewares.VerifyToken())
	{
		userRouters.GET("/profile", userController.Profile)
		userRouters.POST("/profile", userController.Update)
	}

	bookRouters := r.Group("api/v1/book", middlewares.VerifyToken())
	{
		bookRouters.POST("/", bookController.Insert)
		bookRouters.GET("/", bookController.All)
		bookRouters.GET("/:id", bookController.FindByID)
		bookRouters.DELETE("/:id", bookController.Delete)
	}

	r.Run()
}
