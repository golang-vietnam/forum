package routes

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
	"github.com/golang-vietnam/forum/resources"
)

type Home struct{}

var (
	userResource = resources.User{}
)

func (h *Home) Index(c *gin.Context) {
	u := models.User{Name: "Nguyen The Nguyen"}
	userResource.Create(u)
	users, _ := userResource.List()
	c.JSON(200, users)
}

//Use for admin
func (h *Home) AdminDashboard(c *gin.Context) {
	ctx := pongo2.Context{
		"hello": "Hello Admin",
	}
	c.HTML(200, "views/admins/dashboard/index.html", ctx)
}
