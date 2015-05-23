package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name        string        `bson:"name" json:"name"`
	Slug        string        `bson:"slug" json:"slug"`
	Description string        `bson:"description" json:"description"`
	Image       string        `bson:"image" json:"image"`
	CreateAt    time.Time     `bson:"create_at" json:"create_at"`
	UpdateAt    time.Time     `bson:"update_at" json:"update_at"`
}
