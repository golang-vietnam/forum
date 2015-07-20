package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/models"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var (
	categoryValidData = &categoryModel{
		Name:        "category",
		Description: "description",
		Slug:        "slug",
		Image:       "image_url",
	}
)

func CloneCategoryModel(category *categoryModel) *categoryModel {
	newCategory := *category
	return &newCategory
}

func TestCategoryApi(t *testing.T) {
	Convey("Test category api", t, func() {
		Reset(func() {
			clearAll()
		})

		Convey("Create user must success!", func() {
			user := CloneUserModel(userValidData)
			response := do_request("POST", userApi, user)
			body := parse_response(response)
			var responseUser models.User
			err := json.Unmarshal(body, &responseUser)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)

			Convey("Login must success", func() {
				userLogin := CloneUserModel(userValidData)
				response := do_request("POST", authApi+"login", userLogin)
				body := parse_response(response)
				var loginSuccess LoginResponse
				err := json.Unmarshal(body, &loginSuccess)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 200)
				So(loginSuccess.Email, ShouldEqual, userLogin.Email)
				So(loginSuccess.ApiKey, ShouldNotBeNil)

				authsHeader := map[string]string{"Authorization": "Bearer " + loginSuccess.ApiKey}

				Convey("Create category with not admin permission should fail", func() {
					category := CloneCategoryModel(categoryValidData)
					response := do_request("POST", categoryApi, category,
						map[string]string{"Authorization": "Bearer " + loginSuccess.ApiKey})
					body := parse_response(response)
					var responseError Error
					err := json.Unmarshal(body, &responseError)
					So(err, ShouldBeNil)
					So(response.StatusCode, ShouldEqual, 403)
					So(responseError.Id, ShouldEqual, "ACCESS_DENIED")
				})
				Convey("Update user to admin must success", func() {
					update_data := models.User{Role: models.Admin}
					update_error := userCollection().UpdateId(responseUser.Id, &update_data)
					So(update_error, ShouldBeNil)

					Convey("Create category with admin permission should successful and status 201", func() {

						category := CloneCategoryModel(categoryValidData)
						response := do_request("POST", categoryApi, category, authsHeader)
						body := parse_response(response)
						var responseCategory categoryModel
						err := json.Unmarshal(body, &responseCategory)
						So(err, ShouldBeNil)
						So(response.StatusCode, ShouldEqual, 201)
						So(responseCategory.Description, ShouldEqual, category.Description)
						So(responseCategory.Name, ShouldEqual, category.Name)
						So(responseCategory.Slug, ShouldEqual, category.Slug)
						So(responseCategory.Image, ShouldEqual, category.Image)

						Convey("Create exist category with admin permission should fail and status 400", func() {
							category := CloneCategoryModel(categoryValidData)
							response := do_request("POST", categoryApi, category, authsHeader)
							body := parse_response(response)
							var responseError Error
							err := json.Unmarshal(body, &responseError)
							So(err, ShouldBeNil)
							So(response.StatusCode, ShouldEqual, 400)
							So(responseError.Id, ShouldEqual, "CATEGORY_EXIST")
						})

						Convey("Get exist category by Id should response data and status 200", func() {
							response := do_request("GET", categoryApi+responseCategory.Id.Hex(), nil)
							var responseCategory2 categoryModel
							err := json.Unmarshal(body, &responseCategory2)
							So(err, ShouldBeNil)
							So(response.StatusCode, ShouldEqual, 200)
							So(responseCategory2.Name, ShouldEqual, responseCategory.Name)
							So(responseCategory2.Slug, ShouldEqual, responseCategory.Slug)
							So(responseCategory2.Description, ShouldEqual, responseCategory.Description)
						})

						Convey("Get all category should success", func() {
							response := do_request("GET", categoryApi, nil)
							body := parse_response(response)
							var responseCategories Categories
							err := json.Unmarshal(body, &responseCategories)
							So(err, ShouldBeNil)
							So(response.StatusCode, ShouldEqual, 200)
							Convey("Should equal in database", func() {
								var categoriesInDB []userModel
								categoryCollection().Find(nil).All(&categoriesInDB)
								So(len(categoriesInDB), ShouldEqual, len(responseCategories.Categories))
							})
						})
					})
					Convey("Create category with empty name should fail and status 400", func() {
						category := CloneCategoryModel(categoryValidData)
						category.Name = ""
						response := do_request("POST", categoryApi, category, authsHeader)
						body := parse_response(response)
						var responseError Error
						err := json.Unmarshal(body, &responseError)
						So(err, ShouldBeNil)
						So(response.StatusCode, ShouldEqual, 400)
						So(responseError.Id, ShouldEqual, "CATEGORY_NAME_REQUIRED")
					})

					Convey("Create category with empty slug should fail and status 400", func() {
						category := CloneCategoryModel(categoryValidData)
						category.Slug = ""
						response := do_request("POST", categoryApi, category, authsHeader)
						body := parse_response(response)
						var responseError Error
						err := json.Unmarshal(body, &responseError)
						So(err, ShouldBeNil)
						So(response.StatusCode, ShouldEqual, 400)
						So(responseError.Id, ShouldEqual, "CATEGORY_SLUG_REQUIRED")
					})
				})
			})

		})

		Convey("Create category with not login should fail", func() {
			category := CloneCategoryModel(categoryValidData)
			response := do_request("POST", categoryApi, category)
			body := parse_response(response)
			var responseError Error
			err := json.Unmarshal(body, &responseError)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 403)
			So(responseError.Id, ShouldEqual, "ACCESS_DENIED")
		})

		Convey("Get not exist category should fail and status 404", func() {
			response := do_request("GET", categoryApi+bson.NewObjectId().Hex())
			body := parse_response(response)
			var responseError Error
			err := json.Unmarshal(body, &responseError)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 404)
			So(responseError.Id, ShouldEqual, "CATEGORY_NOT_FOUND")
		})

		Convey("Get invalid category Id should fail and status 400", func() {
			response := do_request("GET", categoryApi+"invalid-id")
			body := parse_response(response)
			var responseError Error
			err := json.Unmarshal(body, &responseError)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseError.Id, ShouldEqual, "ID_INVALID")
		})

	})
}
