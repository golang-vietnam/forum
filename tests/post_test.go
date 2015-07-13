package tests

import (
	"github.com/golang-vietnam/forum/database"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var (
	postValidData = &postModel{
		Email:    "ntnguyen@ubisen.com",
		Password: "golang",
		Name:     "nguyen the nguyen",
		Role:     1,

		// Category Category      `bson:"category" json:"category"`
		// Parent   *Post         `bson:"parent" json:"parent"`
		// User     User          `bson:"user" json:"user"`
		// Title    string        `bson:"title" json:"title"`
		// Content  string        `bson:"content" json:"content"`
		// Quote    string        `bson:"quote" json:"quote"`
		// Image    string        `bson:"image" json:"image"`
		// View     int           `bson:"view" json:"view"`
		// Reply    int           `bson:"reply" json:"reply"`
		// Status   int           `bson:"status" json:"status"`
		// LastPost *Post         `bson:"last_post" json:"last_post"`
		// Sticky   int           `bson:"sticky" json:"sticky"`
		// CreateAt time.Time     `bson:"create_at" json:"create_at"`
		// UpdateAt time.Time     `bson:"update_at" json:"update_at"`
	}
)
