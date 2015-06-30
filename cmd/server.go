package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/handlers"
	"github.com/golang-vietnam/forum/middleware"
	"runtime"
	"strconv"
)

func Server() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if _, err := database.InitDb(); err != nil {
		panic(err)
	}

	app := gin.New()
	app.Use(func(c *gin.Context) {
		c.Set(config.SecretKey, config.GetSecret())
		c.Next()
	})
	if config.GetEnv() != config.EnvProduction {
		app.Use(gin.Logger())
		app.Use(gin.Recovery())
	} else {
		app.Use(middleware.Recovery())
	}

	app.Use(middleware.ErrorHandler())
	app.Static("/public", "./public")
	//Set up api v1
	setupApiV1(app)

	env := config.GetEnvValue()
	app.Run(env.Server.Host + ":" + strconv.Itoa(env.Server.Port))
}

func setupApiV1(app *gin.Engine) {
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
	list := []gin.HandlerFunc{userHandler.Create}
	userGroup := v1Group.Group("/user")
	{
		userGroup.GET("/:userId", userHandler.Detail)
		userGroup.POST("/", list...)
	}

	//Post
	postHandler := handlers.NewPostHandler()
	postGroup := v1Group.Group("/post")
	{
		postGroup.GET("/", postHandler.Index)
		postGroup.POST("/", postHandler.Create)
		postGroup.GET("/:id", postHandler.GetById)
	}

	//Auth
	authHandler := handlers.NewAuthHandler()
	authGroup := v1Group.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}
}
