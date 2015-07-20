package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/log"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type postHandlerInterface interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetById(c *gin.Context)
	ListPaging(c *gin.Context)
	ListPagingByCategory(c *gin.Context)
}

func NewPostHandler() postHandlerInterface {
	return &postHandler{}
}

type postHandler struct {
}

func (p *postHandler) Index(c *gin.Context) {
	c.String(200, "Post will be here")
}

func (p *postHandler) Create(c *gin.Context) {
	var post models.Post
	if err := c.Bind(&post); err != nil {
		errors := postResource.ParseError(err)
		if len(errors) > 0 {
			c.Error(errors[0])
			return
		}
	}

	if err := postResource.Create(&post); err != nil {
		log.LogError(c.Request, err, "Error in Create Post", logger)
		c.AbortWithError(500, err)
		return
	}
	c.JSON(201, post)
}

func (p *postHandler) Update(c *gin.Context) {
	var post models.Post
	if err := c.Bind(&post); err != nil {
		errors := postResource.ParseError(err)
		if len(errors) > 0 {
			c.Error(errors[0])
			return
		}
	}

	if err := postResource.Update(&post); err != nil {
		log.LogError(c.Request, err, "Error in Create Post", logger)
		c.AbortWithError(500, err)
		return
	}
	c.JSON(201, post)
}

func (p *postHandler) GetById(c *gin.Context) {
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
		if err != mgo.ErrNotFound {
			log.LogError(c.Request, err, "Error in GetById Post", logger)
			c.AbortWithError(500, err)
			return
		}
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, post)
}

func (p *postHandler) ListPaging(c *gin.Context) {
	var posts []models.Post
	var err error
	pageParam := c.DefaultQuery("page", "1")

	pageIndex, err := strconv.Atoi(pageParam)
	if err != nil {
		pageIndex = 1
	}

	if posts, err = postResource.ListPaging(pageIndex, ITEMS_PER_PAGE); err != nil {
		if err != mgo.ErrNotFound {
			log.LogError(c.Request, err, "Error in ListPaging Post", logger)
			c.AbortWithError(500, err)
			return
		}
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, posts)
}

func (p *postHandler) ListPagingByCategory(c *gin.Context) {
	var posts []models.Post
	var err error
	categoryParam := c.Query("category")
	if categoryParam == "" || !bson.IsObjectIdHex(categoryParam) {
		p.ListPaging(c)
		return
	}
	pageParam := c.DefaultQuery("page", "1")
	pageIndex, err := strconv.Atoi(pageParam)
	if err != nil {
		pageIndex = 1
	}

	categoryId := bson.ObjectIdHex(categoryParam)

	if posts, err = postResource.ListPagingByCategory(categoryId, pageIndex, ITEMS_PER_PAGE); err != nil {
		if err != mgo.ErrNotFound {
			log.LogError(c.Request, err, "Error in ListPagingByCategory Post", logger)
			c.AbortWithError(500, err)
			return
		}
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, posts)
}
