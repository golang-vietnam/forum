package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
	"net/http"
)

type User struct {
}

func (u *User) Index(c *gin.Context) {
	c.String(http.StatusOK, "User page")
}
func (u *User) Create(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return
	}
	if err := userResource.Create(&user); err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(201, user)
}

// Use for Admin page
func (u *User) AdminAllUser(c *gin.Context) {
	c.String(http.StatusOK, "Admin all user page")
}
