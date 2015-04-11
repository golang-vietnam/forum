package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/golang-vietnam/forum/routes"
	"github.com/jinzhu/gorm"
)

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.HTMLRender = middleware.NewPongoRender()

	sqlConnection := "yourdbusername:yourdbpassword@tcp(127.0.0.1:3306)/golang-vietnam?parseTime=True"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
		return
	}
	db.SingularTable(true)

	r.Use(middleware.Gorm("db", db))
	r.NotFound404(routes.Error404)

	homeRouter := &routes.Home{}
	homeGroup := r.Group("/")
	{
		homeGroup.GET("/", homeRouter.Index)
	}

	userRouter := &routes.User{}
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", userRouter.Index)
	}

	r.Run(":8080")
}
