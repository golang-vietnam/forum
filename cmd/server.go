package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/controllers"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/middleware"
	"runtime"
)

func Server() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if _, err := database.InitDb(); err != nil {
		panic(err)
	}
	app := gin.New()
	if config.GetEnv() != config.ENV_PRODUCTION {
		app.Use(gin.Logger())
		app.Use(gin.Recovery())
	} else {
		app.Use(middleware.Recovery())
	}

	app.Use(middleware.ErrorHandler())
	app.Static("/public", "./public")
	//Set up api v1
	setupApiV1(app)

	app.Run(config.GetServer("host") + ":" + config.GetServer("port"))
}

func setupApiV1(app *gin.Engine) {
	//Home
	homeController := controllers.NewHomeController()
	v1Group := app.Group("/v1")
	{
		v1Group.GET("/", homeController.Index)
	}
	apiErrorController := controllers.NewErrorController()
	apiErrorGroup := v1Group.Group("/errors")
	{
		apiErrorGroup.GET("/", apiErrorController.List)
		apiErrorGroup.GET("/:errorId", apiErrorController.GetById)
	}
	//User
	userController := controllers.NewUserController()
	list := []gin.HandlerFunc{userController.Create}
	userGroup := v1Group.Group("/user")
	{
		userGroup.GET("/", userController.Detail)
		userGroup.POST("/", list...)
	}

	//Post
	postController := controllers.NewPostController()
	postGroup := v1Group.Group("/post")
	{
		postGroup.GET("/", postController.Index)
		postGroup.POST("/", postController.Create)
		postGroup.GET("/:id", postController.GetById)
	}

	//Auth
	authController := controllers.NewAuthController()
	authGroup := v1Group.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
	}
}
