package cmd

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/golang-vietnam/forum/routes"
	"github.com/jinzhu/gorm"
)

func Server() {
	app := gin.Default()
	app.Static("/public", "./public")
	app.HTMLRender = middleware.NewPongoRender()

	sqlConnection := "yourdbusername:yourdbpassword@tcp(127.0.0.1:3306)/golangvietnam?parseTime=True"
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

	app.Run(":8080")
}
