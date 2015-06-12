package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/helpers"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/golang-vietnam/forum/routes"
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
	homeRouter := routes.Home{}
	v1Group := app.Group("/v1")
	{
		v1Group.GET("/", homeRouter.Index)
	}

	//User
	userRouter := &routes.User{}
	list := []gin.HandlerFunc{userRouter.Create}
	userGroup := v1Group.Group("/user")
	{
		userGroup.GET("/", userRouter.Detail)
		userGroup.POST("/", list...)
	}

	//Post
	postRouter := &routes.Post{}
	postGroup := v1Group.Group("/post")
	{
		postGroup.GET("/", postRouter.Index)
		postGroup.POST("/", postRouter.Create)
		postGroup.GET("/:id", postRouter.GetById)
	}

	//Auth
	authRouter := &routes.Auth{}
	authGroup := v1Group.Group("/auth")
	{
		authGroup.GET("/", authRouter.Provider)
		authGroup.GET("/callback", authRouter.CallBack)
	}
}
