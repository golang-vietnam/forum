package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/helpers"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/golang-vietnam/forum/routes"
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
	app := gin.Default()
	app.Use(middleware.Goth(authsConf))
	app.Use(middleware.ErrorHandler())
	app.Static("/public", "./public")
	app.HTMLRender = helpers.NewPongRender()

	homeRouter := routes.Home{}
	homeGroup := app.Group("/v1")
	{
		homeGroup.GET("/", homeRouter.Index)
	}

	userRouter := &routes.User{}
	userGroup := app.Group("v1/user")
	{
		userGroup.GET("/", userRouter.Index)
		userGroup.POST("/", userRouter.Create)
	}
	authRouter := &routes.Auth{}
	authGroup := app.Group("v1/auth")
	{
		authGroup.GET("/", authRouter.Provider)
		authGroup.GET("/callback", authRouter.CallBack)
	}
	app.Run(config.GetServer("host") + ":" + config.GetServer("port"))

}
