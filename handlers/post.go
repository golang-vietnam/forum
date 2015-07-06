package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type postHandlerInterface interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	GetById(c *gin.Context)
	ListPaging(c *gin.Context)
	ListPagingByCategory(c *gin.Context)
}

const (
	ITEMS_PER_PAGE = 15
)

func NewPostHandler() postHandlerInterface {
	return &postController{}
}

type postHandler struct {
}

func (p *postHandler) Index(c *gin.Context) {
	c.String(200, "Post will be here")
}

/*
   [WIP] Not complete yet, just some demo for test purpose
*/
func (p *postHandler) Create(c *gin.Context) {
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
		panic(err)
	}

	c.JSON(200, post)
}

func (p *postController) ListPaging(c *gin.Context) {
	var posts []models.Post
	var err error
	pageParam := c.DefaultQuery("page", "1")

	pageIndex, err := strconv.Atoi(pageParam)
	if err != nil {
		pageIndex = 1
	}

	posts = postResource.ListPaging(pageIndex, ITEMS_PER_PAGE)
	c.JSON(200, posts)
}

func (p *postController) ListPagingByCategory(c *gin.Context) {
	var posts []models.Post
	var err error
	pageParam := c.DefaultQuery("page", "1")
	categoryParam := c.DefaultQuery("category", "1")

	pageIndex, err := strconv.Atoi(pageParam)
	if err != nil {
		pageIndex = 1
	}

	pageIndex, err := strconv.Atoi(pageParam)
	if err != nil {
		pageIndex = 1
	}

	posts = postResource.ListPaging(pageIndex, ITEMS_PER_PAGE)
	c.JSON(200, posts)
}
