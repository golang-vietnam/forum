package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
)

type userHandlerInterface interface {
	Detail(c *gin.Context)
	Create(c *gin.Context)
	Edit(c *gin.Context)
}

func NewUserHandler() userHandlerInterface {
	return &userHandler{}
}

type userHandler struct {
}

func (u *userHandler) Detail(c *gin.Context) {
	userId := c.Param("userId")
	user, err := userResource.GetById(userId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, user)
}

/**

	TODO:
	- Parse request with validate
	- Create user with valid request
	- Reponse new user

**/

func (u *userHandler) Create(c *gin.Context) {
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

func (u *userHandler) Edit(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		errors := userResource.ParseError(err)
		if len(errors) > 0 {
			c.Error(errors[0])
			return
		}
	}
	userId := c.Param("userId")
	if err := userResource.Edit(userId, &user); err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, user)
}
