package resources

import (
	// "errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
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

func (r *ResourceUser) Create(u *models.User) *apiErrors.Error {
	u.Id = bson.NewObjectId()
	if err := collection("user").Insert(u); err != nil {
		if mgo.IsDup(err) {
			return &apiErrors.USER_EXIST
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
func (r *ResourceUser) ParseError(err error) *apiErrors.Error {
	if errs, ok := err.(*validator.StructErrors); ok {
		for _, v := range errs.Errors {
			switch v.Field {
			case "Email":
				switch v.Tag {
				case "required":
					return &apiErrors.USER_EMAIL_REQUIRER
				case "email":
					return &apiErrors.USER_EMAIL_INVALID
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
