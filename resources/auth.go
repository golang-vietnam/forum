package resources

import (
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
)

type ResourceAuthInterface interface {
	Login(email string, password string) (*models.User, error)
}

type resourceAuth struct {
}

var (
	newResourceUser = NewResourceUser()
)

func NewResourceAuth() ResourceAuthInterface {
	return &resourceAuth{}
}

func (r *resourceAuth) Login(email string, password string) (*models.User, error) {
	user := newResourceUser.GetByEmail(email)
	if user == nil {
		return nil, &apiErrors.AUTH_EMAIL_INVALID
	}
	if !newResourceUser.IsMatchPassword(user.Password, password) {
		return nil, &apiErrors.AUTH_PASSWORD_INVALID
	}
	return user, nil
}
