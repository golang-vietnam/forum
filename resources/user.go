package resources

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResourceUserInterface interface {
	List() ([]models.User, error)
	GetById(id bson.ObjectId) (models.User, error)
	Create(u *models.User) error
	RemoveById(id bson.ObjectId) error
	Validate(u *models.User) []error
	ParseError(err error) []error
	IsMatchPassword(hashedPassword string, password string) bool
	HashPassword(password string) string
}

func NewResourceUser() ResourceUserInterface {
	return &ResourceUser{}
}

type ResourceUser struct {
}

const userColName = "user"

func (r *ResourceUser) List() ([]models.User, error) {
	var users []models.User
	err := collection(userColName).Find(nil).All(&users)
	return users, err
}

func (r *ResourceUser) GetById(id bson.ObjectId) (models.User, error) {
	var user models.User
	err := collection(userColName).FindId(id).One(&user)
	return user, err
}

/**

	TODO:
	- Check password not be nil
	- Hash password
	- Gen new MongoObjectId for Id
	- Set role user to 0
	- Insert user to db
	- Check error return is exist

**/

func (r *ResourceUser) Create(u *models.User) error {
	if u.Password == "" {
		return &apiErrors.USER_PASSWORD_REQUIRED
	}
	u.Password = r.HashPassword(u.Password)
	u.Id = bson.NewObjectId()
	u.Role = 0
	if err := collection(userColName).Insert(u); err != nil {
		if mgo.IsDup(err) {
			return &apiErrors.USER_EXIST
		}
		panic(err)
	}
	return nil
}

func (r *ResourceUser) RemoveById(id bson.ObjectId) error {
	return collection(userColName).RemoveId(id)
}

func (r *ResourceUser) Validate(u *models.User) []error {
	if err := binding.Validate(u); err != nil {
		return r.ParseError(err)
	}
	return []error{}
}
func (r *ResourceUser) ParseError(err error) []error {
	var errors []error
	if errs, ok := err.(*validator.StructErrors); ok {
		for _, v := range errs.Errors {
			switch v.Field {
			case "Email":
				switch v.Tag {
				case "required":
					errors = append(errors, &apiErrors.USER_EMAIL_REQUIRED)
				case "email":
					errors = append(errors, &apiErrors.USER_EMAIL_INVALID)
				case "max":
					errors = append(errors, &apiErrors.USER_EMAIL_MAX)
				case "min":
					errors = append(errors, &apiErrors.USER_EMAIL_MIN)
				}
			case "Password":
				switch v.Tag {
				case "required":
					errors = append(errors, &apiErrors.USER_PASSWORD_REQUIRED)
				default:
					return nil
				}
			case "Role":
				switch v.Tag {
				case "max":
					errors = append(errors, &apiErrors.USER_ROLE_MAX)
				case "min":
					errors = append(errors, &apiErrors.USER_ROLE_MIN)
				}
			}
		}
	} else {
		panic("Can not parse error")
	}

	return errors
}
func (r *ResourceUser) HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
func (r *ResourceUser) IsMatchPassword(hashedPassword string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
