package resources

import (
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2/bson"
)

type resourcePostInterface interface {
	ListAll() ([]models.Post, error)
	ListPaging(page int, perPage int) ([]models.Post, error)
	ListPagingByCategory(id bson.ObjectId, page int, perPage int) ([]models.Post, error)
	GetById(id bson.ObjectId) (models.Post, error)
	Create(u *models.Post) error
	RemoveById(id bson.ObjectId) error
}

func NewResourcePost() resourcePostInterface {
	return &resourcePost{}
}

type resourcePost struct {
}

const postColName = models.PostColName

func (r *resourcePost) ListAll() ([]models.Post, error) {
	var posts []models.Post
	err := collection(postColName).Find(nil).All(&posts)
	return posts, err
}

//Please make sure page >= 0 && perPage >= 0
func (r *resourcePost) ListPaging(page int, perPage int) ([]models.Post, error) {
	var posts []models.Post
	err := collection(postColName).Find(nil).Limit(perPage).Skip(perPage * page).All(&posts)
	return posts, err
}

//Please make sure page >= 0 && perPage >= 0
func (r *resourcePost) ListPagingByCategory(categoryId bson.ObjectId, page int, perPage int) ([]models.Post, error) {
	var posts []models.Post
	err := collection(postColName).Find(bson.M{"category._id": categoryId}).Limit(perPage).Skip(perPage * page).All(&posts)
	return posts, err
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

func (r *resourcePost) RemoveById(id bson.ObjectId) error {
	err := collection(postColName).RemoveId(id)
	return err
}
