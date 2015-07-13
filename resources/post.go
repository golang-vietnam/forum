package resources

import (
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2"
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
	if err := collection(postColName).Find(nil).All(&posts); err != nil {
		if err == mgo.ErrNotFound {
			return nil, apiErrors.ThrowError(apiErrors.PostNotFound)
		}
		panic(err)
	}
	return posts, err
}

//Please make sure page >= 0 && perPage >= 0
func (r *resourcePost) ListPaging(page int, perPage int) ([]models.Post, error) {
	var posts []models.Post
	if err := collection(postColName).Find(nil).Limit(perPage).Skip(perPage * page).All(&posts); err != nil {
		panic(err)
	}
	return posts, err
}

//Please make sure page >= 0 && perPage >= 0
func (r *resourcePost) ListPagingByCategory(categoryId bson.ObjectId, page int, perPage int) ([]models.Post, error) {
	var posts []models.Post
	if err := collection(postColName).Find(bson.M{"category._id": categoryId}).Limit(perPage).Skip(perPage * page).All(&posts); err != nil {
		panic(err)
	}
	return posts, err
}

func (r *resourcePost) GetById(id bson.ObjectId) (models.Post, error) {
	var post models.Post
	if err := collection(postColName).FindId(id).One(&post); err != nil {
		if err == mgo.ErrNotFound {
			return nil, apiErrors.ThrowError(apiErrors.PostNotFound)
		}
		panic(err)
	}
	return post, err
}

func (r *resourcePost) Create(p *models.Post) error {
	p.Id = bson.NewObjectId()
	if err := collection(postColName).Insert(p); err != nil {
		if mgo.IsDup(err) {
			return apiErrors.ThrowError(apiErrors.PostExist)
		}
		panic(err)
	}
	return err
}

func (r *resourcePost) RemoveById(id bson.ObjectId) error {
	if err := collection(postColName).RemoveId(id); err != nil {
		if err == mgo.ErrNotFound {
			return apiErrors.ThrowError(apiErrors.PostNotFound)
		}
		panic(err)
	}
	return err
}
