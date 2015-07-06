package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type LoginResponse struct {
	models.User `json:"user"`
	ApiKey      string `json:"api-key"`
}

func TestAuthen(t *testing.T) {

	Convey("POST Login", t, func() {
		Reset(func() {
			database.ClearAll()
		})

		Convey("Register new account must successful!", func() {
			user := CloneUserModel(userValidData)
			response := do_request("POST", userApi, user)
			body := parse_response(response)
			var responseUser models.User
			err := json.Unmarshal(body, &responseUser)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)

			Convey("Login with correct account should successful!", func() {
				user := CloneUserModel(userValidData)
				response := do_request("POST", authApi+"login", user)
				body := parse_response(response)
				var loginSuccess LoginResponse
				err := json.Unmarshal(body, &loginSuccess)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 200)
				So(loginSuccess.Email, ShouldEqual, user.Email)
				So(loginSuccess.ApiKey, ShouldNotBeNil)
				Println(loginSuccess.Password)
			})

		})
		Convey("Login with not have account should fail", func() {
			user := CloneUserModel(userValidData)
			user.Email = "nothaveuser@email.com"
			response := do_request("POST", authApi+"login", user)
			body := parse_response(response)

			var loginFail Error
			err := json.Unmarshal(body, &loginFail)
			So(err, ShouldBeNil)
		})

	})
}
