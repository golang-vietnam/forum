package routes

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func Error404(c *gin.Context) {
	ctx := pongo2.Context{}
	c.HTML(200, "views/errors/404.html", ctx)
}
