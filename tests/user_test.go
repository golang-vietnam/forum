package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/database"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	userValidData = &userModel{
		Email:    "ntnguyen@ubisen.com",
		Password: "golang",
		Name:     "nguyen the nguyen",
		Role:     1,
	}
	userInvalidEmailData = &userModel{
		Email:    "invalidemail",
		Name:     "invalidEmail",
		Password: "invalidEmail",
	}
	userInvalidEmailEmptyData = &userModel{
		Name:     "Empty Email",
		Password: "Empty Email",
	}
	userInvalidPasswordEmptyData = &userModel{
		Name:  "Ntn",
		Email: "ntnguyen@ubisen.com",
	}
)

func CloneUserModel(user *userModel) *userModel {
	newUser := *user
	return &newUser
}

func TestUser(t *testing.T) {

	database.ClearAllUser()

	Convey("GET user", t, func() {
		Convey("Get not exist user shoud return not found error", func() {
			So(1, ShouldEqual, 1)
		})
	})

	Convey("POST create user", t, func() {

		Convey("Create not exist user should response status 201 and correct user data.", func() {
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

			var userInDb userModel
			database.Collection(UserColName).FindId(responseUser.Id).One(&userInDb)
			So(userInDb.Id, ShouldEqual, responseUser.Id)

			Convey("Create exist user should response status 400 and exist message", func() {
				response := do_request("POST", userApi, userValidData)
				body := parse_response(response)
				var responseError Error
				err := json.Unmarshal(body, &responseError)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 400)
				So(responseError.Id, ShouldEqual, "USER_EXIST")
				So(responseError.Message, ShouldEqual, "This user has been exist!")
			})
		})

		Convey("Create with invalid email should return status 400 and email invalid message", func() {
			response := do_request("POST", userApi, userInvalidEmailData)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Id, ShouldEqual, "USER_EMAIL_INVALID")
			So(responseData.Message, ShouldEqual, "Email invalid")
		})
		Convey("Create with empty email should return status 400 and email required message", func() {
			response := do_request("POST", userApi, userInvalidEmailEmptyData)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Id, ShouldEqual, "USER_EMAIL_REQUIRED")
			So(responseData.Message, ShouldEqual, "Email is required")
		})
		Convey("Create with empty password should return status 400 and password required message", func() {
			response := do_request("POST", userApi, userInvalidPasswordEmptyData)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Id, ShouldEqual, "USER_PASSWORD_REQUIRED")
			So(responseData.Message, ShouldEqual, "Password is required")
		})
	})
}
