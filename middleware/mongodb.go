package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"gopkg.in/mgo.v2"
	"time"
)

func Mongo() gin.HandlerFunc {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.GetDB("host")},
		Timeout:  60 * time.Second,
		Database: config.GetDB("name"),
		Username: config.GetDB("user"),
		Password: config.GetDB("password"),
	}
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.Set("db", mongoSession)
		defer mongoSession.Close()
		c.Next()
	}
}

func MongoSession(c *gin.Context) *mgo.Session {
	return c.MustGet("db").(*mgo.Session)
}
