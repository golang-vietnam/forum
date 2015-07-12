package resources

import (
	"github.com/golang-vietnam/forum/database"
	"gopkg.in/mgo.v2"
)

var (
	newResourceUser = NewResourceUser()
)

func collection(c string) *mgo.Collection {
	return database.Collection(c)
}
