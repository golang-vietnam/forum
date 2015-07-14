package tests

import (
	"encoding/json"
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
			clearAll()
		})

		Convey("Register new account must successful!", func() {
			user := CloneUserModel(userValidData)
			response := do_request("POST", userApi, user)
			body := parse_response(response)
			var responseUser models.User
			err := json.Unmarshal(body, &responseUser)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)

			Convey("Update exist user not login should fail", func() {
				user2 := CloneUserModel(userValidData)
				user2.Name = "New Name"
				user2.Id = responseUser.Id
				response := do_request("PUT", userApi+responseUser.Id.Hex(), user2)
				body := parse_response(response)
				var responseError Error
				err := json.Unmarshal(body, &responseError)
				So(err, ShouldBeNil)
				So(responseError.Id, ShouldEqual, "ACCESS_DENIED")
				So(response.StatusCode, ShouldEqual, 403)
			})

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

				Convey("Update exist user should success", func() {
					user := CloneUserModel(userValidData)
					user.Name = "New Name"
					user.Id = newObjectId()
					response := do_request("PUT", userApi+responseUser.Id.Hex(), user,
						map[string]string{"Authorization": "Bearer " + loginSuccess.ApiKey})
					body := parse_response(response)
					var userRes userModel
					err := json.Unmarshal(body, &userRes)
					So(err, ShouldBeNil)
					So(response.StatusCode, ShouldEqual, 200)
					So(userRes.Name, ShouldEqual, user.Name)
					So(userRes.Id.Hex(), ShouldNotEqual, user.Id.Hex())
				})
				Convey("Create new user must successful", func() {
					user2 := CloneUserModel(userValidData)
					user2.Email = "newemail@abc.com"
					response2 := do_request("POST", userApi, user2)
					body := parse_response(response2)
					var responseUser2 userModel
					err := json.Unmarshal(body, &responseUser2)
					So(err, ShouldBeNil)
					So(response2.StatusCode, ShouldEqual, 201)
					So(responseUser2.Email, ShouldEqual, user2.Email)

					Convey("Should not update another user", func() {
						user3 := CloneUserModel(userValidData)
						user3.Email = "newemail3@abc.com"
						response := do_request("PUT", userApi+responseUser2.Id.Hex(), user3,
							map[string]string{"Authorization": "Bearer " + loginSuccess.ApiKey})
						body := parse_response(response)
						var responseError Error
						err := json.Unmarshal(body, &responseError)
						So(err, ShouldBeNil)
						So(response.StatusCode, ShouldEqual, 403)
						So(responseError.Id, ShouldEqual, "ACCESS_DENIED")
					})
				})

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

		Convey("Login with invalid email should fail", func() {
			user := CloneUserModel(userValidData)
			user.Email = "notemail"
			response := do_request("POST", authApi+"login", user)
			body := parse_response(response)

			var loginFail Error
			err := json.Unmarshal(body, &loginFail)
			So(err, ShouldBeNil)
		})

	})
}
