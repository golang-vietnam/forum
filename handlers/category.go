package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/log"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2"
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
		errors := categoryResource.ParseError(err)
		if len(errors) > 0 {
			c.Error(errors[0])
			return
		}
	}
	if err := categoryResource.Create(&category); err != nil {
		log.LogError(c.Request, err, "Error in Create Category", logger)
		c.Error(err)
		return
	}
	c.JSON(201, category)
}

func (p *categoryHandler) Update(c *gin.Context) {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		errors := categoryResource.ParseError(err)
		if len(errors) > 0 {
			c.Error(errors[0])
			return
		}
	}

	if err := categoryResource.Update(&category); err != nil {
		if err != mgo.ErrNotFound {
			log.LogError(c.Request, err, "Error in Update Category", logger)
			c.AbortWithError(500, err)
			return
		}
		c.AbortWithError(404, err)
		return
	}
	c.JSON(201, category)
}

func (p *categoryHandler) GetById(c *gin.Context) {
	category := c.MustGet("categoryData")
	c.JSON(200, category)
}

func (p *categoryHandler) GetAll(c *gin.Context) {
	categories := categoryResource.GetAll()
	c.JSON(200, gin.H{"categories": categories})
}
