package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	NormalUser  = 0
	EditorUser  = 1
	UserColName = "users"
)

type User struct {
	Id           bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name         string        `bson:"name" json:"name"`
	Email        string        `bson:"email" json:"email" binding:"required,max=50,min=3,email"`
	Password     string        `bson:"password" json:"password" binding:"required"`
	LastActivate time.Time     `bson:"last_activate" json:"last_activate"`
	Image        string        `bson:"image" json:"image"`
	Status       string        `bson:"status" json:"status"`
	Role         int           `bson:"role" json:"role" binding:"max=1,min=0"`
	CreateAt     time.Time     `bson:"create_at" json:"create_at"`
	UpdateAt     time.Time     `bson:"update_at" json:"update_at"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
