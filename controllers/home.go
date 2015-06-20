package controllers

import (
	// "github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

type homeControllerInterface interface {
	Index(c *gin.Context)
}

func NewHomeController() homeControllerInterface {
	return &home{}
}

type home struct{}

func (h *home) Index(c *gin.Context) {
	c.JSON(200, "Hello forum API")
}

//Use for admin
// func (h *home) AdminDashboard(c *gin.Context) {

// 	ctx := pongo2.Context{
// 		"hello": "Hello Admin",
// 	}
// 	c.HTML(200, "views/admins/dashboard/index.html", ctx)
// }
