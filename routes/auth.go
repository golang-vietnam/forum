package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthInterface interface {
	CallBack(c *gin.Context)
	Provider(c *gin.Context)
}

type Auth struct{}

/**

	TODO:
	- Get data from oauth2 provider
	- Update or create new user
	- Response user data with request is logined
	- Notes: This route will not be use for api.

**/

func (a *Auth) CallBack(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}
	c.JSON(200, user)
}

/**

	TODO:
	- Begin login oauth2 to provider
	- Notes: This route will not be use for api.

**/

func (a *Auth) Provider(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}
