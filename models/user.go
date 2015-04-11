package models

type User struct {
	Id   int
	name string
}

func (u *User) Create() (user, err) {
	return nil, nil
}
