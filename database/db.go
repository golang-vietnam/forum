package database

import (
	"github.com/golang-vietnam/forum/config"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	mongoSession *mgo.Session
	dataBase     string
)

func InitDb() (*mgo.Session, error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.GetDB("host")},
		Timeout:  60 * time.Second,
		Database: config.GetDB("name"),
		Username: config.GetDB("user"),
		Password: config.GetDB("password"),
	}
	if mongoSession != nil {
		panic("mongodb is connected!")
	}
	var err error
	if mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo); err != nil {
		return mongoSession, err
	}
	dataBase = config.GetDB("name")
	return mongoSession, Index()
}

func Index() error {
	// Collection("user").E
	return Collection("user").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}

func Collection(c string) *mgo.Collection {
	return mongoSession.DB(dataBase).C(c)
}

func CloseDb() {
	if mongoSession != nil {
		mongoSession.Close()
	}
}

func ClearAllUser() {
	if _, err := Collection("user").RemoveAll(nil); err != nil {
		panic(err)
	}
}
