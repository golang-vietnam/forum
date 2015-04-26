package routes

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/models"
)

type Home struct{}

func (h *Home) Index(c *gin.Context) {
	var users2 []models.User
	user := database.Collection(c, "user")
	err := user.Find(nil).All(&users2)
	if err != nil {
		panic(err)
	}
	c.JSON(200, users2)
}

//Use for admin
func (h *Home) AdminDashboard(c *gin.Context) {
	ctx := pongo2.Context{
		"hello": "Hello Admin",
	}
	c.HTML(200, "views/admins/dashboard/index.html", ctx)
}
