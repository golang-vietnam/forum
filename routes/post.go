package routes

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	m "github.com/golang-vietnam/forum/models"
)

type Post struct {
}

func (p *Post) Index(c *gin.Context) {
	c.String(200, "Post will be here")
}

/*
   [WIP] Not complete yet, just some demo for test purpose
*/
func (p *Post) Create(c *gin.Context) {
	var post m.Post
	if err := c.Bind(&post); err != nil {
		c.AbortWithError(400, err)
	}
	if err := postResource.Create(&post); err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(201, post)
}

func (p *Post) GetById(c *gin.Context) {
	var post m.Post
	var err error
	id := c.Params.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		c.AbortWithStatus(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	if post, err = postResource.GetById(oid); err != nil {
		panic(err)
	}

	c.JSON(200, post)
}
