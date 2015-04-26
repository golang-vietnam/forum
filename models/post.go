package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Category Category      `bson:"category"`
	Parent   *Post         `bson:"parent"`
	User     User          `bson:"user"`
	Title    string        `bson:"title"`
	Content  string        `bson:"content"`
	Qoute    string        `bson:"qoute"`
	Image    string        `bson:"image"`
	View     int           `bson:"view"`
	Reply    int           `bson:"reply"`
	Status   int           `bson:"status"`
	LastPost *Post         `bson:"last_post"`
	Sticky   int           `bson:"sticky"`
	CreateAt time.Time     `bson:"create_at"`
	UpdateAt time.Time     `bson:"update_at"`
}
