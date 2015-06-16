package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
)

type UserControllerInterface interface {
	Detail(c *gin.Context)
	Create(c *gin.Context)
}

func NewUserController() UserControllerInterface {
	return &userController{}
}

type userController struct {
}

func (u *userController) Detail(c *gin.Context) {
	c.String(200, "User page")
}

/**

	TODO:
	- Parse request with validate
	- Create user with valid request
	- Reponse new user

**/

func (u *userController) Create(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		errors := userResource.ParseError(err)
		if len(errors) > 0 {
			c.Error(errors[0])
			return
		}

	}
	if err := userResource.Create(&user); err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(201, user)
}
