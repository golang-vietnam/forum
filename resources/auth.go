package resources

import (
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
)

type resourceAuthInterface interface {
	Login(email string, password string) (*models.User, error)
}

type resourceAuth struct {
}

var (
	newResourceUser = NewResourceUser()
)

func NewResourceAuth() resourceAuthInterface {
	return &resourceAuth{}
}

func (r *resourceAuth) Login(email string, password string) (*models.User, error) {
	user, err := newResourceUser.GetByEmail(email)
	if err != nil {
		return nil, apiErrors.ThrowError(apiErrors.AuthEmailInvalid)
	}
	if !newResourceUser.IsMatchPassword(user.Password, password) {
		return nil, apiErrors.ThrowError(apiErrors.AuthPasswordInValid)
	}
	return user, nil
}
