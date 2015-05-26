package resources

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
	if err := collection("user").Insert(u); err != nil {
		if mgo.IsDup(err) {
			return errors.New("This user has been exist!")
		}
		panic(err)
	}
	return nil
}

func (r *ResourceUser) RemoveById(id bson.ObjectId) error {
	return collection("user").RemoveId(id)
}

func (r *ResourceUser) Validate(u *models.User) error {
	if err := binding.Validate(u); err != nil {
		return r.ParseError(err)
	}
	return nil
}
func (r *ResourceUser) ParseError(err error) error {
	if errs, ok := err.(*validator.StructErrors); ok {
		for _, v := range errs.Errors {
			switch v.Field {
			case "Email":
				switch v.Tag {
				case "required":
					return errors.New("Email is required")
				case "email":
					return errors.New("Email invalid")
				default:
					return nil
				}
			default:
				return nil
			}
		}
	} else {
		panic("Can not parse error")
	}

	return nil
}
func ClearAllUser() {
	if _, err := collection("user").RemoveAll(nil); err != nil {
		panic(err)
	}
}
