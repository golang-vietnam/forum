package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/utils"
	"github.com/golang-vietnam/forum/models"
)

type userHandlerInterface interface {
	Index(c *gin.Context)
	Detail(c *gin.Context)
	Create(c *gin.Context)
	Edit(c *gin.Context)
}

func NewUserHandler() userHandlerInterface {
	return &userHandler{}
}

type userHandler struct {
}

func (u *userHandler) Index(c *gin.Context) {
	userResource.GenerateAvatar("yolo")
	c.JSON(200, "Yolo")
}

func (u *userHandler) Detail(c *gin.Context) {
	user := c.MustGet("userData")
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

	currentUser := utils.MustGetCurrentUser(c)

	if err := c.Bind(&user); err != nil {
		errors := userResource.ParseError(err)
		if len(errors) > 0 {
			c.Error(errors[0])
			return
		}
	}
	userId := c.Param("userId")

	if currentUser.Role != models.Admin {
		user.Role = models.NormalUser
	}
	if currentUser.Role == models.NormalUser {
		user.DeleteAt = nil
	}

	if err := userResource.Edit(userId, &user); err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(200, user)
}
