package cmd

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/golang-vietnam/forum/routes"
	"github.com/jinzhu/gorm"
)

func Server() {
	app := gin.Default()
	app.Static("/public", "./public")
	app.HTMLRender = middleware.NewPongoRender()

	sqlConnection := config.GetDB("user") + ":" + config.GetDB("password") +
		"@tcp(" + config.GetDB("host") + ":" + config.GetDB("port") + ")/" + config.GetDB("name") + "?parseTime=True"

	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
		return
	}
	db.SingularTable(true)

	app.Use(middleware.Gorm("db", db))
	app.NotFound404(routes.Error404)

	homeRouter := &routes.Home{}
	homeGroup := app.Group("/")
	{
		homeGroup.GET("/", homeRouter.Index)
	}

	userRouter := &routes.User{}
	userGroup := app.Group("/user")
	{
		userGroup.GET("/", userRouter.Index)
	}
	app.Run(config.GetServer("host") + ":" + config.GetServer("port"))
}
