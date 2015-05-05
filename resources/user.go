package resources

import (
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2/bson"
)

type ResourceUserInterface interface {
	List() ([]models.User, error)
	GetById(id bson.ObjectId) (models.User, error)
	Create(u *models.User) error
	RemoveById(id bson.ObjectId) error
	Validate(u models.User) (bool, error)
}

type ResourceUser struct {
}

func (r *ResourceUser) List() ([]models.User, error) {
	var users []models.User
	err := collection("user").Find(nil).All(&users)
	return users, err
}

func (r *ResourceUser) GetById(id bson.ObjectId) (models.User, error) {
	var user models.User
	err := collection("user").FindId(id).One(&user)
	return user, err
}

func (r *ResourceUser) Create(u *models.User) error {
	u.Id = bson.NewObjectId()
	return collection("user").Insert(u)
}

func (r *ResourceUser) RemoveById(id bson.ObjectId) error {
	return collection("user").RemoveId(id)
}

func (r *ResourceUser) Validate(u models.User) (bool, error) {
	return false, nil
}
