package routes

import (
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/golang-vietnam/forum/models"
)

type Home struct{}

func (h *Home) Index(c *gin.Context) {
	mongoSession := middleware.MongoSession(c)
	db := mongoSession.DB(config.GetDB("name"))
	user := db.C("user")
	user.Insert(&models.User{Name: "Balabala"})
	var users []models.User
	err := db.C("user").Find(nil).All(&users)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, users)
}

//Use for admin
func (h *Home) AdminDashboard(c *gin.Context) {
	ctx := pongo2.Context{
		"hello": "Hello Admin",
	}
	c.HTML(200, "views/admins/dashboard/index.html", ctx)
}
