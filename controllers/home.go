package controllers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

type Home struct{}

func (h *Home) Index(c *gin.Context) {
	c.JSON(200, "Hello forum API")
}

//Use for admin
func (h *Home) AdminDashboard(c *gin.Context) {

	ctx := pongo2.Context{
		"hello": "Hello Admin",
	}
	c.HTML(200, "views/admins/dashboard/index.html", ctx)
}
