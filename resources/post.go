package resources

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResourcePostInterface interface {
	ListAll() ([]models.Post, error)
	ListPaging() ([]models.Post, error)
	GetById(id bson.ObjectId) (models.Post, error)
	Create(u *models.Post) error
	RemoveById(id bson.ObjectId) error
}

func NewResourcePost() ResourcePostInterface {
	return &ResourcePost{}
}

type ResourcePost struct {
}

const colName = "post"

func (r *ResourcePost) ListAll() ([]models.Post, error) {
	var posts []models.Post
	err := collection(colName).Find(nil).All(&posts)
	return posts, err
}

func (r *ResourcePost) ListPaging(limit int) ([]models.Post, error) {
	var posts []models.Post
	err := collection(colName).Find(nil).Limit(limit).All(&posts)
	return posts, err
}

func (r *ResourcePost) GetById(id bson.ObjectId) (models.Post, error) {
	var post models.Post
	err := collection(colName).FindId(id).One(&post)
	return post, err
}

func (r *ResourcePost) Create(p *models.Post) error {
	p.Id = bson.NewObjectId()
	err := collection(colName).Insert(p)
	return err
}

func (r *ResourcePost) RemoveById(id bson.ObjectId) error {
	return collection(colName).RemoveId(id)
}
