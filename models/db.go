package models

import (
	// "fmt"
	"github.com/golang-vietnam/forum/config"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	mongoSession *mgo.Session
)

func InitDb() error {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.GetDB("host")},
		Timeout:  60 * time.Second,
		Database: config.GetDB("name"),
		Username: config.GetDB("user"),
		Password: config.GetDB("password"),
	}
	var err error
	mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	_ = mongoSession
	return err
}

func collection(c string) *mgo.Collection {
	return mongoSession.DB(config.GetDB("name")).C(c)
}
