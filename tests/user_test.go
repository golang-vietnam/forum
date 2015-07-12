package tests

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var (
	userValidData = &userModel{
		Email:    "ntnguyen@ubisen.com",
		Password: "golang",
		Name:     "nguyen the nguyen",
		Role:     1,
	}
)

func CloneUserModel(user *userModel) *userModel {
	newUser := *user
	return &newUser
}

func TestUserNotLoginApi(t *testing.T) {

	Convey("Test user when not login API", t, func() {
		Reset(func() {
			clearAll()
		})

		Convey("Create with invalid email should return status 400 and email invalid message", func() {
			user := CloneUserModel(userValidData)
			user.Email = "invalidEmail"
			response := do_request("POST", userApi, user)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Id, ShouldEqual, "USER_EMAIL_INVALID")
			So(responseData.Message, ShouldEqual, "Email invalid")

		})
		Convey("Create with empty email should return status 400 and email required message", func() {
			user := CloneUserModel(userValidData)
			user.Email = ""
			response := do_request("POST", userApi, user)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Id, ShouldEqual, "USER_EMAIL_REQUIRED")
			So(responseData.Message, ShouldEqual, "Email is required")
		})
		Convey("Create with empty password should return status 400 and password required message", func() {
			user := CloneUserModel(userValidData)
			user.Password = ""
			response := do_request("POST", userApi, user)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Id, ShouldEqual, "USER_PASSWORD_REQUIRED")
			So(responseData.Message, ShouldEqual, "Password is required")
		})

		Convey("Create not exist user should response status 201 and correct user data.", func() {

			Convey("Should success", func() {
				user := CloneUserModel(userValidData)
				response := do_request("POST", userApi, user)
				body := parse_response(response)
				var responseUser userModel
				err := json.Unmarshal(body, &responseUser)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 201)
				So(responseUser.Email, ShouldEqual, user.Email)
				So(responseUser.Name, ShouldEqual, user.Name)
				So(responseUser.Role, ShouldEqual, 0)

				Convey("Should in database", func() {
					var userInDb userModel
					userCollection().FindId(responseUser.Id).One(&userInDb)
					So(userInDb.Id, ShouldEqual, responseUser.Id)

				})

				Convey("Get exist user should response status 200 and user info", func() {
					response := do_request("GET", userApi+responseUser.Id.Hex())
					body := parse_response(response)
					var user userModel
					err := json.Unmarshal(body, &user)
					So(err, ShouldBeNil)
					So(response.StatusCode, ShouldEqual, 200)

				})
				Convey("Get invalid id should fail", func() {
					user := CloneUserModel(userValidData)
					response := do_request("GET", userApi+"helloId", user)
					body := parse_response(response)
					var responseError Error
					err := json.Unmarshal(body, &responseError)
					So(err, ShouldBeNil)
					So(responseError.Id, ShouldEqual, "ID_INVALID")
					So(response.StatusCode, ShouldEqual, 400)
				})

				Convey("Create exist user should response status 400 and exist message", func() {
					user := CloneUserModel(userValidData)
					response := do_request("POST", userApi, user)
					body := parse_response(response)
					var responseError Error
					err := json.Unmarshal(body, &responseError)
					So(err, ShouldBeNil)
					So(response.StatusCode, ShouldEqual, 400)
					So(responseError.Id, ShouldEqual, "USER_EXIST")
					So(responseError.Message, ShouldEqual, "This user has been exist!")
				})

			})

		})
		Convey("Update not exist user should fail", func() {
			user := CloneUserModel(userValidData)
			user.Name = "New Name2"
			response := do_request("PUT", userApi+bson.NewObjectId().Hex(), user)
			body := parse_response(response)
			var responseError Error
			err := json.Unmarshal(body, &responseError)
			So(err, ShouldBeNil)
			So(responseError.Id, ShouldEqual, "USER_NOT_FOUND")
			So(response.StatusCode, ShouldEqual, 404)
		})

		Convey("Update invalid id should fail", func() {
			user := CloneUserModel(userValidData)
			user.Name = "New Name"
			response := do_request("PUT", userApi+"helloId", user)
			body := parse_response(response)
			var responseError Error
			err := json.Unmarshal(body, &responseError)
			So(err, ShouldBeNil)
			So(responseError.Id, ShouldEqual, "ID_INVALID")
			So(response.StatusCode, ShouldEqual, 400)
		})
	})
}
