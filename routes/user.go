package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
)

type User struct {
}

func (u *User) Detail(c *gin.Context) {
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
