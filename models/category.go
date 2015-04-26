package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Slug        string        `bson:"slug"`
	Description string        `bson:"description"`
	Image       string        `bson:"image"`
	CreateAt    time.Time     `bson:"create_at"`
	UpdateAt    time.Time     `bson:"update_at"`
}
