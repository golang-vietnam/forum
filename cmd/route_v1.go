package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/handlers"
	"github.com/golang-vietnam/forum/middleware"
)

func routeV1(app *gin.Engine) {
	loads := middleware.NewLoads()
	auths := middleware.NewAuthMiddleware()

	//Home
	homeHandler := handlers.NewHomeHandler()
	v1Group := app.Group("/v1")
	{
		v1Group.GET("/", homeHandler.Index)
	}
	apiErrorHandler := handlers.NewErrorHandler()
	apiErrorGroup := v1Group.Group("/errors")
	{
		apiErrorGroup.GET("/", apiErrorHandler.List)
		apiErrorGroup.GET("/:errorId", apiErrorHandler.GetById)
	}
	//User
	userHandler := handlers.NewUserHandler()
	userEdit := []gin.HandlerFunc{
		loads.LoadUserById(),
		auths.RequireLogin(),
		auths.UserHasAuthorization(),
		userHandler.Edit,
	}
	userGroup := v1Group.Group("/users")
	{
		userGroup.GET("/:userId", loads.LoadUserById(), userHandler.Detail)
		userGroup.PUT("/:userId", userEdit...)
		userGroup.POST("/", userHandler.Create)
	}

	//Post
	postHandler := handlers.NewPostHandler()
	postGroup := v1Group.Group("/posts")
	{
		postGroup.GET("/", postHandler.Index)
		postGroup.POST("/", postHandler.Create)
		postGroup.GET("/:id", postHandler.GetById)
	}

	//Category
	categoryHandler := handlers.NewCategoryHandler()
	categoryGroup := v1Group.Group("/categories")
	{
		categoryGroup.GET("/", categoryHandler.GetAll)
		categoryGroup.POST("/", categoryHandler.Create)
		categoryGroup.PUT("/", categoryHandler.Update)
		categoryGroup.GET("/:categoryId", loads.LoadCategoryById(), categoryHandler.GetById)
	}

	//Auth
	authHandler := handlers.NewAuthHandler()
	authGroup := v1Group.Group("/auths")
	{
		authGroup.POST("/login", authHandler.Login)
	}
}
