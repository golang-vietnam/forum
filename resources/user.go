package resources

import (
	"github.com/golang-vietnam/forum/models"
)

type UserInterface interface {
	List() []models.User
	Create(um models.User)
}

type User struct {
}

func (u *User) List() ([]models.User, error) {
	var users []models.User
	err := collection("user").Find(nil).All(&users)
	return users, err
}
func (u *User) Create(um models.User) {
	collection("user").Insert(&um)
}
