package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/helpers"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/golang-vietnam/forum/resources"
	"github.com/golang-vietnam/forum/routes"
	"runtime"
)

func Server() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	err := resources.InitDb()
	if err != nil {
		panic(err)
	}
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(middleware.Recovery())
	app.Use(middleware.ErrorHandler())
	app.Static("/public", "./public")
	app.HTMLRender = helpers.NewPongoRender()
	app.NotFound404(routes.Error404)

	homeRouter := routes.Home{}
	homeGroup := app.Group("/")
	{
		homeGroup.GET("/", homeRouter.Index)
	}

	userRouter := &routes.User{}
	userGroup := app.Group("/user")
	{
		userGroup.GET("/", userRouter.Index)
	}
	adminGroup := app.Group("/admin")
	{
		adminGroup.GET("/", homeRouter.AdminDashboard)
		userAdmin := adminGroup.Group("/user")
		{
			userAdmin.GET("/", userRouter.AdminAllUser)
		}
	}
	app.Run(config.GetServer("host") + ":" + config.GetServer("port"))
}
