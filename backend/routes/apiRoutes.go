package routes

import (
	"github.com/Arkar27/gin-bulletinboard/backend/controllers"
	authDao "github.com/Arkar27/gin-bulletinboard/backend/dao/authDao"
	postDao "github.com/Arkar27/gin-bulletinboard/backend/dao/postDao"
	userDao "github.com/Arkar27/gin-bulletinboard/backend/dao/userDao"
	"github.com/Arkar27/gin-bulletinboard/backend/initializers"
	middleware "github.com/Arkar27/gin-bulletinboard/backend/middleware"
	services "github.com/Arkar27/gin-bulletinboard/backend/services"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(route *gin.Engine) {

	loginDao := authDao.NewAuthDao(initializers.DB)
	loginService := services.NewLoginService(loginDao)
	loginController := controllers.NewLoginController(loginService)

	userDao := userDao.NewUserDao(initializers.DB)
	userService := services.NewUserService(userDao)
	userController := controllers.NewUserController(userService)

	postDao := postDao.NewPostDao(initializers.DB)
	postService := services.NewPostService(postDao)
	postController := controllers.NewPostController(postService)

	apiRouter := route.Group("/api")

	// for swagger route "/api/swagger/index.html"
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// login
	apiRouter.POST("/login", loginController.Login)

	// user CRUD routes
	userRoute := apiRouter.Group("/users")
	{
		userRoute.GET("/", middleware.AuthMiddleware(), userController.GetUserList)
		userRoute.GET("/:id", middleware.AuthMiddleware(), userController.GetUserById)
		userRoute.POST("/", middleware.AuthMiddleware(), userController.CreateUser)
		userRoute.PUT("/:id", middleware.AuthMiddleware(), userController.UpdateUser)
		userRoute.DELETE("/:id", middleware.AuthMiddleware(), userController.DeleteUser)
	}

	// post CRUD routes
	postRoute := apiRouter.Group("/posts")
	{
		postRoute.GET("/", middleware.AuthMiddleware(), postController.GetPostList)
		postRoute.GET("/:id", middleware.AuthMiddleware(), postController.GetPostById)
		postRoute.POST("/", middleware.AuthMiddleware(), postController.CreatePost)
		postRoute.PUT("/:id", middleware.AuthMiddleware(), postController.UpdatePost)
		postRoute.DELETE("/:id", middleware.AuthMiddleware(), postController.DeletePost)
	}
}
