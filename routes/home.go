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
	ctx := pongo2.Context{
		"hello": "Hello text from route and use in child layout",
	}
	mongoSession := middleware.MongoSession(c)
	db := mongoSession.DB(config.GetDB("name"))
	user := db.C("user")
	user.Insert(&models.User{Name: "Balabala"})
	var users []models.User
	err := db.C("user").Find(nil).All(&users)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].Id)
	}

	c.HTML(200, "views/sites/home/index.html", ctx)
}
