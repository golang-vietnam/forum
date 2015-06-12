package resources

import (
	// "github.com/gin-gonic/gin/binding"
	// "github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	// "gopkg.in/bluesuncorp/validator.v5"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResourcePostInterface interface {
	ListAll() []models.Post
	ListPaging(page int, perPage int) []models.Post
	ListPagingByCategory(id bson.ObjectId, page int, perPage int) []models.Post
	GetById(id bson.ObjectId) (models.Post, error)
	Create(u *models.Post) error
	RemoveById(id bson.ObjectId)
}

func NewResourcePost() ResourcePostInterface {
	return &resourcePost{}
}

type resourcePost struct {
}

const postColName = "posts"

func (r *resourcePost) ListAll() []models.Post {
	var posts []models.Post
	if err := collection(postColName).Find(nil).All(&posts); err != nil {
		panic(err)
	}
	return posts
}

func (r *resourcePost) ListPaging(page int, perPage int) []models.Post {
	if page < 0 || perPage < 0 {
		panic("list paging post param invalid")
	}
	var posts []models.Post
	if err := collection(postColName).Find(nil).Limit(perPage).Skip(perPage * page).All(&posts); err != nil {
		panic(err)
	}
	return posts
}

func (r *resourcePost) ListPagingByCategory(categoryId bson.ObjectId, page int, perPage int) []models.Post {
	if page < 0 || perPage < 0 {
		panic("list paging post param invalid")
	}
	var posts []models.Post
	//Not tested this query yet
	if err := collection(postColName).Find(bson.M{"category._id": categoryId}).Limit(perPage).Skip(perPage * page).All(&posts); err != nil {
		panic(err)
	}
	return posts
}

func (r *resourcePost) GetById(id bson.ObjectId) (models.Post, error) {
	var post models.Post
	err := collection(postColName).FindId(id).One(&post)
	return post, err
}

func (r *resourcePost) Create(p *models.Post) error {
	p.Id = bson.NewObjectId()
	err := collection(postColName).Insert(p)
	return err
}

func (r *resourcePost) RemoveById(id bson.ObjectId) {
	if err := collection(postColName).RemoveId(id); err != nil {
		panic(err)
	}
}
