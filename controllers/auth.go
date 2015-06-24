package controllers

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	"time"
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
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	apiKey, err := token.SignedString([]byte(config.GetSecret()))
	if err != nil {
		c.Error(apiErrors.ThrowError(apiErrors.ServerError))
		return
	}
	// Remove password
	user.Password = ""

	c.JSON(200, gin.H{
		"user":    user,
		"api-key": apiKey,
	})
}
