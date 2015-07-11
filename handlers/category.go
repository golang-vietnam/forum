package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2/bson"
)

type categoryHandlerInterface interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
}

func NewCategoryHandler() categoryHandlerInterface {
	return &categoryHandler{}
}

type categoryHandler struct {
}

func (p *categoryHandler) Index(c *gin.Context) {
	c.String(200, "Category will be here")
}

func (p *categoryHandler) Create(c *gin.Context) {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		c.AbortWithError(400, err)
		return
	}
	if err := categoryResource.Create(&category); err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(201, category)
}

func (p *categoryHandler) Update(c *gin.Context) {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		c.AbortWithError(400, err)
		return
	}

	if err := categoryResource.Update(&category); err != nil {
		c.AbortWithError(400, err)
		return
	}
	c.JSON(201, category)
}

func (p *categoryHandler) GetById(c *gin.Context) {
	var category models.Category
	var err error
	id := c.Params.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		c.AbortWithStatus(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	if category, err = categoryResource.GetById(oid); err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, category)
}

func (p *categoryHandler) GetAll(c *gin.Context) {
	var categories []models.Category
	var err error

	if categories, err = categoryResource.GetAll(); err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, categories)
}
