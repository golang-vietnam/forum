package resources

import (
	"errors"
	m "github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

func (r ResourceUser) Create(u *m.User) error {
	u.Id = bson.NewObjectId()
	if err := collection("user").Insert(u); err != nil {
		if mgo.IsDup(err) {
			return errors.New("This user has been exist!")
		}
		panic(err)
	}
	return nil
}

func (r ResourceUser) RemoveById(id bson.ObjectId) error {
	return collection("user").RemoveId(id)
}

func (r ResourceUser) Validate(u m.User) (bool, error) {
	return false, nil
}
func ClearAllUser() {
	if _, err := collection("user").RemoveAll(nil); err != nil {
		panic(err)
	}
}
