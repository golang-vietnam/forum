package resources

import (
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"gopkg.in/mgo.v2"
)

func collection(c string) *mgo.Collection {
	return database.Collection(c)
}

func newApiError(apiError *apiErrors.Error) *apiErrors.Error {
	return apiErrors.CloneError(apiError)
}
