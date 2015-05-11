package routes

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	m "github.com/golang-vietnam/forum/models"
	r "github.com/golang-vietnam/forum/resources"
)

type Home struct{}

var (
	ru = r.ResourceUser{}
)

func (h *Home) Index(c *gin.Context) {
	u := m.User{Name: "Nguyen The Nguyen"}
	u.Email = "ntnguyen@ubisen.com"
	err := ru.Create(&u)
	if !err.IsNil() {
		return
	}
	c.JSON(200, u)
}

//Use for admin
func (h *Home) AdminDashboard(c *gin.Context) {
	ctx := pongo2.Context{
		"hello": "Hello Admin",
	}
	c.HTML(200, "views/admins/dashboard/index.html", ctx)
}
