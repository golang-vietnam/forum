package routes

import (
	_ "fmt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
)

type Home struct{}

func (h *Home) Index(c *gin.Context) {
	ctx := pongo2.Context{
		"hello": "Hello text from route and use in child layout",
	}
	//db := c.Get("db")
	c.HTML(200, "views/sites/home/index.html", ctx)
}
