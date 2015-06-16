package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2/bson"
)

type PostControllerInterface interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	GetById(c *gin.Context)
}

func NewPostController() PostControllerInterface {
	return &postController{}
}

type postController struct {
}

func (p *postController) Index(c *gin.Context) {
	c.String(200, "Post will be here")
}

/*
   [WIP] Not complete yet, just some demo for test purpose
*/
func (p *postController) Create(c *gin.Context) {
	var post models.Post
	if err := c.Bind(&post); err != nil {
		c.AbortWithError(400, err)
	}
	if err := postResource.Create(&post); err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(201, post)
}

func (p *postController) GetById(c *gin.Context) {
	var post models.Post
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
