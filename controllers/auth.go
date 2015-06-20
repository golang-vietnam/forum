package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
)

type authControllerInterface interface {
	Login(c *gin.Context)
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
