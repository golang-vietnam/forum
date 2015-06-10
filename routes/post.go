package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
)

type Post struct {
}

func (p *Post) Index(c *gin.Context) {
	c.String(200, "Post will be here")
}

func (p *Post) Create(c *gin.Context) {
	//Do some awesome things here
}
