package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAuthen(t *testing.T) {
	database.ClearAllUser()
	Convey("POST Login", t, func() {
		Convey("Register new account must successful!", func() {
			user := &models.User{Email: "ntnguyen@ubisen.com", Name: "Nguyen The Nguyen", Password: "golang", Role: 1}
			response := do_request("POST", userApi, user)
			body := parse_response(response)
			var responseUser models.User
			err := json.Unmarshal(body, &responseUser)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)

			Convey("Login with correct account should successful!", func() {
				response := do_request("POST", authApi+"login", &models.UserLogin{Email: "ntnguyen@ubisen.com", Password: "golang"})
				body := parse_response(response)
				var responseUser models.User
				err := json.Unmarshal(body, &responseUser)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 200)
				So(responseUser.Email, ShouldEqual, "ntnguyen@ubisen.com")
			})
		})

	})
}
