package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
	"github.com/markbates/goth/gothic"
)

type authControllerInterface interface {
	Login(c *gin.Context)
	CallBack(c *gin.Context)
	Provider(c *gin.Context)
}

func NewAuthController() authControllerInterface {
	return &authController{}
}

type authController struct{}

/**

	TODO:
	- Get email and password from post request
	- Find user

**/
func (a *authController) Login(c *gin.Context) {
	var userLogin models.UserLogin
	if err := c.Bind(&userLogin); err != nil {
		panic(err)
	}
	user, err := authResource.Login(userLogin.Email, userLogin.Password)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, user)
}

/**

	TODO:
	- Get data from oauth2 provider
	- Update or create new user
	- Response user data with request is logined
	- Notes: This route will not be use for api.

**/

func (a *authController) CallBack(c *gin.Context) {
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

func (a *authController) Provider(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}
