package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	categoryPostValidData = models.Category{
		Name:        "category",
		Description: "description",
		Slug:        "slug",
		Image:       "image_url",
	}

	userPostValidData = models.User{
		Email:    "ntnguyen@ubisen.com",
		Password: "golang",
		Name:     "nguyen the nguyen",
		Role:     1,
	}

	postValidParent = &models.Post{
		Category: categoryPostValidData,
		Parent:   nil,
		User:     userPostValidData,
		Title:    "Title",
		Content:  "Title",
		Quote:    "Quote",
		Image:    "image",
		View:     0,
		Reply:    0,
		Status:   0,
		LastPost: nil,
		Sticky:   1,
	}

	postValidChild = &postModel{
		Category: categoryPostValidData,
		Parent:   postValidParent,
		User:     userPostValidData,
		Title:    "Title",
		Content:  "Title",
		Quote:    "Quote",
		Image:    "image",
		View:     0,
		Reply:    0,
		Status:   0,
		LastPost: nil,
		Sticky:   1,
	}
)

func ClonePostModel(post *models.Post) *models.Post {
	newPost := *post
	return &newPost
}

func TestPostApi(t *testing.T) {
	Convey("Test post api", t, func() {
		Reset(func() {
			clearAll()
		})

		Convey("Create Post should successful and status 201", func() {
			post := ClonePostModel(postValidParent)
			response := do_request("POST", postApi, post)
			body := parse_response(response)
			var responsePost postModel
			err := json.Unmarshal(body, &responsePost)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)
			So(responsePost.Content, ShouldEqual, post.Content)
			So(responsePost.Quote, ShouldEqual, post.Quote)
			So(responsePost.Title, ShouldEqual, post.Title)
			// So(responsePost.Category, ShouldEqual, post.Category)
			So(responsePost.Image, ShouldEqual, post.Image)
			So(responsePost.Reply, ShouldEqual, post.Reply)
			So(responsePost.Status, ShouldEqual, post.Status)
			So(responsePost.Sticky, ShouldEqual, post.Sticky)
			// So(responsePost.User, ShouldEqual, post.User)
			So(responsePost.View, ShouldEqual, post.View)

			Convey("Get an exists Post by Id should response data and status 200", func() {
				response := do_request("GET", postApi+responsePost.Id.Hex(), nil)
				var responsePost2 postModel
				err := json.Unmarshal(body, &responsePost2)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 200)
				So(responsePost2.Content, ShouldEqual, responsePost.Content)
				So(responsePost2.Quote, ShouldEqual, responsePost.Quote)
				So(responsePost2.Title, ShouldEqual, responsePost.Title)
				// So(responsePost2.Category, ShouldEqual, responsePost.Category)
				So(responsePost2.Image, ShouldEqual, responsePost.Image)
				So(responsePost2.Reply, ShouldEqual, responsePost.Reply)
				So(responsePost2.Status, ShouldEqual, responsePost.Status)
				So(responsePost2.Sticky, ShouldEqual, responsePost.Sticky)
				// So(responsePost2.User, ShouldEqual, responsePost.User)
				So(responsePost2.View, ShouldEqual, responsePost.View)
				// So(responsePost2.Name, ShouldEqual, responsePost.Name)
				// So(responsePost2.Slug, ShouldEqual, responsePost.Slug)
				// So(responsePost2.Description, ShouldEqual, responsePost.Description)
			})
			// Convey("Get all Posts should success", func() {
			// 	response := do_request("GET", postApi, nil)
			// 	body := parse_response(response)
			// 	var responsePosts []models.Post
			// 	err := json.Unmarshal(body, &responsePosts)
			// 	So(err, ShouldBeNil)
			// 	So(response.StatusCode, ShouldEqual, 200)
			// 	Convey("Should equal in database", func() {
			// 		var postsInDB []userModel
			// 		postCollection().Find(nil).All(&postsInDB)
			// 		So(len(postsInDB), ShouldEqual, len(responsePosts))
			// 	})
			// })
		})

		// Convey("Create post with empty name should fail and status 400", func() {
		// 	post := ClonePostModel(postValidParent)
		// 	post.Name = ""
		// 	response := do_request("POST", postApi, post)
		// 	body := parse_response(response)
		// 	var responseError Error
		// 	err := json.Unmarshal(body, &responseError)
		// 	So(err, ShouldBeNil)
		// 	So(response.StatusCode, ShouldEqual, 400)
		// 	So(responseError.Id, ShouldEqual, "CATEGORY_NAME_REQUIRED")
		// })
		// Convey("Create post with empty slug should fail and status 400", func() {
		// 	post := ClonePostModel(postValidParent)
		// 	post.Slug = ""
		// 	response := do_request("POST", postApi, post)
		// 	body := parse_response(response)
		// 	var responseError Error
		// 	err := json.Unmarshal(body, &responseError)
		// 	So(err, ShouldBeNil)
		// 	So(response.StatusCode, ShouldEqual, 400)
		// 	So(responseError.Id, ShouldEqual, "CATEGORY_SLUG_REQUIRED")
		// })
		// Convey("Get not exist post should fail and status 404", func() {
		// 	response := do_request("GET", postApi+bson.NewObjectId().Hex())
		// 	body := parse_response(response)
		// 	var responseError Error
		// 	err := json.Unmarshal(body, &responseError)
		// 	So(err, ShouldBeNil)
		// 	So(response.StatusCode, ShouldEqual, 404)
		// 	So(responseError.Id, ShouldEqual, "CATEGORY_NOT_FOUND")
		// })

		// Convey("Get invalid post Id should fail and status 400", func() {
		// 	response := do_request("GET", postApi+"invalid-id")
		// 	body := parse_response(response)
		// 	var responseError Error
		// 	err := json.Unmarshal(body, &responseError)
		// 	So(err, ShouldBeNil)
		// 	So(response.StatusCode, ShouldEqual, 400)
		// 	So(responseError.Id, ShouldEqual, "ID_INVALID")
		// })

	})
}
