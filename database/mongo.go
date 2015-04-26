package database

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"gopkg.in/mgo.v2"
)

func Collection(c *gin.Context, cname string) *mgo.Collection {
	return c.MustGet("db").(*mgo.Session).DB(config.GetDB("name")).C(cname)
}
