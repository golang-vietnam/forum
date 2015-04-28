package models

import (
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

// func (u *User) Save() {
// 	collection("user").Insert(u)
// }
