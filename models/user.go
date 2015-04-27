package models

import (
	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Email        string        `bson:"email"`
	Password     string        `bson:"password"`
	LastActivate time.Time     `bson:"last_activate"`
	Image        string        `bson:"image"`
	Status       string        `bson:"status"`
	Role         int           `bson:"role"`
	Category     []Category    `bson:"category"`
	CreateAt     time.Time     `bson:"create_at"`
	UpdateAt     time.Time     `bson:"update_at"`
}

func (u *User) All() []User {
	var users []User
	collection("user").Find(nil).All(&users)
	return users
}
func (u *User) Save() {
	collection("user").Insert(u)
}
