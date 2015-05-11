package resources

import (
	h "github.com/golang-vietnam/forum/helpers"
	m "github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// type Resource interface {
// 	List() ([]m.User, error)
// 	GetById(id bson.ObjectId) (m.User, error)
// 	Create(r *m.User) h.Error
// 	RemoveById(id bson.ObjectId) error
// 	Validate(u m.User) (bool, error)
// }

type ResourceUser struct {
}

func (r ResourceUser) List() ([]m.User, error) {
	var users []m.User
	err := collection("user").Find(nil).All(&users)
	return users, err
}

func (r ResourceUser) GetById(id bson.ObjectId) (m.User, error) {
	var user m.User
	err := collection("user").FindId(id).One(&user)
	return user, err
}

func (r ResourceUser) Create(u *m.User) h.Error {
	u.Id = bson.NewObjectId()
	err := collection("user").Insert(u)
	if err != nil {
		if mgo.IsDup(err) {
			badRequest := h.ErrBadRequest
			badRequest.Detail = "This account has been exist!"
			return badRequest
		}
		panic(err)
	}
	return h.Error{}
}

func (r ResourceUser) RemoveById(id bson.ObjectId) error {
	return collection("user").RemoveId(id)
}

func (r ResourceUser) Validate(u m.User) (bool, error) {
	return false, nil
}
