package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
)

type User struct {
}

func (u *User) Index(c *gin.Context) {
	c.String(200, "User page")
}

/**

	TODO:
	- Parse request with validate
	- Create user with valid request
	- Reponse new user

**/

func (u *User) Create(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.Error(userResource.ParseError(err))
		return
	}
	if err := userResource.Create(&user); err != nil {

		c.AbortWithError(400, err)
		return
	}
	c.JSON(201, user)
}
