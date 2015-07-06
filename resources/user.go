package resources

import (
	// "fmt"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const userColName = models.UserColName

type resourceUserInterface interface {
	ListAll() []*models.User
	GetById(id string) (*models.User, error)
	GetByEmail(email string) *models.User
	Create(u *models.User) error
	Edit(id string, u *models.User) error
	RemoveById(id bson.ObjectId)
	ParseError(err error) []error
	IsMatchPassword(hashedPassword string, password string) bool
	HashPassword(password string) string
}

func NewResourceUser() resourceUserInterface {
	return &resourceUser{}
}

type resourceUser struct {
}

func (r *resourceUser) ListAll() []*models.User {
	var users []*models.User
	if err := collection(userColName).Find(nil).All(&users); err != nil {
		panic(err)
	}
	return users
}

func (r *resourceUser) GetById(id string) (*models.User, error) {

	if !bson.IsObjectIdHex(id) {
		return nil, apiErrors.ThrowError(apiErrors.UserIdInValid)
	}

	var user models.User
	if err := collection(userColName).FindId(bson.ObjectIdHex(id)).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			return nil, apiErrors.ThrowError(apiErrors.UserNotFound)
		}
		panic(err)
	}
	return &user, nil
}

func (r *resourceUser) GetByEmail(email string) *models.User {
	var user models.User
	if err := collection(userColName).Find(bson.M{"email": email}).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return &user
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
		return apiErrors.ThrowError(apiErrors.UserPasswordRequired)
	}
	u.Password = r.HashPassword(u.Password)
	u.Id = bson.NewObjectId()
	u.Role = models.NormalUser
	if err := collection(userColName).Insert(u); err != nil {
		if mgo.IsDup(err) {
			return apiErrors.ThrowError(apiErrors.UserExist)
		}
		panic(err)
	}

	return nil
}

func (r *resourceUser) Edit(id string, u *models.User) error {
	if !bson.IsObjectIdHex(id) {
		return apiErrors.ThrowError(apiErrors.UserIdInValid)
	}
	if err := collection(userColName).UpdateId(bson.ObjectIdHex(id), u); err != nil {
		if err == mgo.ErrNotFound {
			return apiErrors.ThrowError(apiErrors.UserNotFound)
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
					errors = append(errors, apiErrors.ThrowError(apiErrors.UserEmailRequied))
				case "email":
					errors = append(errors, apiErrors.ThrowError(apiErrors.UserEmailInvalid))
				case "max":
					errors = append(errors, apiErrors.ThrowError(apiErrors.UserEmailMax))
				case "min":
					errors = append(errors, apiErrors.ThrowError(apiErrors.UserEmailMin))
				}
			case "Password":
				switch v.Tag {
				case "required":
					errors = append(errors, apiErrors.ThrowError(apiErrors.UserPasswordRequired))
				}
			case "Role":
				switch v.Tag {
				case "max":
					errors = append(errors, apiErrors.ThrowError(apiErrors.UserRoleMax))
				case "min":
					errors = append(errors, apiErrors.ThrowError(apiErrors.UserRoleMin))
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
