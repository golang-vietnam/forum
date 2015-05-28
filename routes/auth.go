package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type Auth struct{}

func (a *Auth) CallBack(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}
	c.JSON(200, user)
}
func (a *Auth) Provider(c *gin.Context) {

	gothic.BeginAuthHandler(c.Writer, c.Request)
}
