package routes

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/middleware"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2"
)

type Home struct{}

var (
	mongoSession *mgo.Session
)

func (h *Home) Index(c *gin.Context) {
	ctx := pongo2.Context{
		"hello": "Hello text from route and use in child layout",
	}
	mongoSession = middleware.MongoSession(c)
	db := mongoSession.DB(config.GetDB("name"))
	user := db.C("user")
	user.Insert(&models.User{Name: "Balabala"})
	c.HTML(200, "views/sites/home/index.html", ctx)
}
