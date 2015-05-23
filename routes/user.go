package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers"
	"github.com/golang-vietnam/forum/models"
	"github.com/golang-vietnam/forum/resources"
	"net/http"
)

type User struct {
}

func (u *User) Index(c *gin.Context) {
	resources.ClearAllUser()
	c.String(http.StatusOK, "User page")
}
func (u *User) Create(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		helpers.AddError(c, helpers.ErrBadRequest)
		return
	}

	if err := userResource.Create(&user); !err.IsNil() {
		helpers.AddError(c, err)
		return
	}

	c.JSON(201, user)
}

// Use for Admin page
func (u *User) AdminAllUser(c *gin.Context) {
	c.String(http.StatusOK, "Admin all user page")
}
