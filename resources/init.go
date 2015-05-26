package resources

import (
	"github.com/golang-vietnam/forum/config"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	mongoSession *mgo.Session
	dataBase     string
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
	if mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo); err != nil {
		return err
	}
	dataBase = config.GetDB("name")
	return index()
}

func index() error {
	return collection("user").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
}

func collection(c string) *mgo.Collection {
	return mongoSession.DB(dataBase).C(c)
}

func CloseDb() {
	if mongoSession != nil {
		mongoSession.Close()
	}
}
