package models

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	Email     string
	password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u *User) Create() (user, err) {
	return nil, nil
}
