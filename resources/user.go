package resources

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResourceUserInterface interface {
	List() ([]models.User, error)
	GetById(id bson.ObjectId) (models.User, error)
	Create(u *models.User) error
	RemoveById(id bson.ObjectId) error
	Validate(u *models.User) error
	ParseError(err error) error
}

func NewResourceUser() ResourceUserInterface {
	return &ResourceUser{}
}

type ResourceUser struct {
}

const colName = "user"

func (r *ResourceUser) List() ([]models.User, error) {
	var users []models.User
	err := collection(colName).Find(nil).All(&users)
	return users, err
}

func (r *ResourceUser) GetById(id bson.ObjectId) (models.User, error) {
	var user models.User
	err := collection(colName).FindId(id).One(&user)
	return user, err
}

func (r *ResourceUser) Create(u *models.User) error {
	u.Id = bson.NewObjectId()
	if err := collection(colName).Insert(u); err != nil {
		if mgo.IsDup(err) {
			return &apiErrors.USER_EXIST
		}
		panic(err)
	}
	return nil
}

func (r *ResourceUser) RemoveById(id bson.ObjectId) error {
	return collection(colName).RemoveId(id)
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
					return &apiErrors.USER_EMAIL_REQUIRED
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
