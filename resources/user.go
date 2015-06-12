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

const userColName = "users"

type ResourceUserInterface interface {
	ListAll() []models.User
	GetById(id bson.ObjectId) models.User
	Create(u *models.User) error
	RemoveById(id bson.ObjectId)
	Validate(u *models.User) []error
	ParseError(err error) []error
	IsMatchPassword(hashedPassword string, password string) bool
	HashPassword(password string) string
}

func NewResourceUser() ResourceUserInterface {
	return &resourceUser{}
}

type resourceUser struct {
}

func (r *resourceUser) ListAll() []models.User {
	var users []models.User
	if err := collection(userColName).Find(nil).All(&users); err != nil {
		panic(err)
	}
	return users
}

func (r *resourceUser) GetById(id bson.ObjectId) models.User {
	var user models.User
	if err := collection(userColName).FindId(id).One(&user); err != nil {
		panic(err)
	}
	return user
}

/**

	TODO:
	- Check password not be nil
	- Hash password
	- Gen new MongoObjectId for Id
	- Set role user to NormalUser for attack
	- Insert user to db
	- Check error return is exist

**/

func (r *resourceUser) Create(u *models.User) error {
	if u.Password == "" {
		return &apiErrors.USER_PASSWORD_REQUIRED
	}
	u.Password = r.HashPassword(u.Password)
	u.Id = bson.NewObjectId()
	u.Role = models.NormalUser
	if err := collection(userColName).Insert(u); err != nil {
		if mgo.IsDup(err) {
			return &apiErrors.USER_EXIST
		}
		panic(err)
	}
	return nil
}

func (r *resourceUser) RemoveById(id bson.ObjectId) {
	if err := collection(userColName).RemoveId(id); err != nil {
		panic(err)
	}
}

func (r *resourceUser) Validate(u *models.User) []error {
	if err := binding.Validate(u); err != nil {
		return r.ParseError(err)
	}
	return []error{}
}

/**

	TODO:
	- Parse validator error to api error

**/

func (r *resourceUser) ParseError(err error) []error {
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
func (r *resourceUser) HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
func (r *resourceUser) IsMatchPassword(hashedPassword string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
