package database

import (
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	mongoSession *mgo.Session
	dataBase     string
)

func InitDb() (*mgo.Session, error) {
	env := config.GetEnvValue()

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{env.Database.Host},
		Timeout:  60 * time.Second,
		Database: env.Database.Name,
		Username: env.Database.User,
		Password: env.Database.Password,
	}
	if mongoSession != nil {
		panic("mongodb is connected!")
	}
	var err error
	if mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo); err != nil {
		return mongoSession, err
	}
	dataBase = env.Database.Name
	return mongoSession, Index()
}

func Index() error {
	if err := Collection(models.UserColName).EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		return err
	}
	if err := Collection(models.CategoryColName).EnsureIndex(mgo.Index{
		Key:    []string{"slug"},
		Unique: true,
	}); err != nil {
		return err
	}
	return nil
}

func Collection(c string) *mgo.Collection {
	return mongoSession.DB(dataBase).C(c)
}

func CloseDb() {
	if mongoSession != nil {
		mongoSession.Close()
	}
}
func ClearAll() {
	ClearAllUser()
	ClearAllCategory()
}
func ClearAllUser() {
	if _, err := Collection(models.UserColName).RemoveAll(nil); err != nil {
		panic(err)
	}
}
func ClearAllCategory() {
	if _, err := Collection(models.CategoryColName).RemoveAll(nil); err != nil {
		panic(err)
	}
}
