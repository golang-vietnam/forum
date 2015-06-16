package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/controllers"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/helpers"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/spf13/viper"
	"runtime"
)

func Server() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if _, err := database.InitDb(); err != nil {
		panic(err)
	}
	var (
		urlAuth   = "http://" + config.GetServer("host") + ":" + config.GetServer("port") + "/v1/auth/"
		authsConf = map[string][]string{
			"facebook": []string{"1578087022454903", "2aff5458c8645a998103d00c99085938", urlAuth + "callback?provider=facebook"},
			// "google":   []string{"", "", urlAuth + "/callback?provider=google"},
			// "github":   []string{"", "", urlAuth + "/callback?provider=github"},
		}
	)
	fmt.Println(urlAuth)
	app := gin.New()
	if viper.Get("env") != "production" {
		app.Use(gin.Logger())
		app.Use(gin.Recovery())
	} else {
		app.Use(middleware.Recovery())
	}

	app.Use(middleware.Goth(authsConf))
	app.Use(middleware.ErrorHandler())
	app.Static("/public", "./public")
	app.HTMLRender = helpers.NewPongRender()

	//Set up api v1
	setupApiV1(app)

	app.Run(config.GetServer("host") + ":" + config.GetServer("port"))
}

func setupApiV1(app *gin.Engine) {
	//Home
	homeController := controllers.Home{}
	v1Group := app.Group("/v1")
	{
		v1Group.GET("/", homeController.Index)
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
		authGroup.GET("/", authController.Provider)
		authGroup.GET("/callback", authController.CallBack)
	}
}
