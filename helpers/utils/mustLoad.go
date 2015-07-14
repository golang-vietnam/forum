package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
)

func MustGetCurrentUser(c *gin.Context) *models.User {
	var user *models.User
	var ok bool
	if data, err := c.Get("currentUser"); !err {
		panic("Must use RequireLogin before")
	} else {
		if user, ok = data.(*models.User); !ok {
			panic("currentUser must *models.User type")
		}
	}
	return user
}

func MustGetUserData(c *gin.Context) *models.User {
	var user *models.User
	var ok bool
	if data, err := c.Get("userData"); !err {
		panic("Must use LoadUserById before")
	} else {
		if user, ok = data.(*models.User); !ok {
			panic("currentUser must *models.User type")
		}
	}
	return user
}
