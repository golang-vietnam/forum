package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/models"
	"github.com/golang-vietnam/forum/resources"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUser(t *testing.T) {
	server := getServer()

	resources.InitDb()
	defer resources.CloseDb()
	resources.ClearAllUser()

	Convey("POST create user", t, func() {

		url := server + "/api/user/"
		Convey("Create not exist user should response status 201 and correct error data.", func() {

			user := &models.User{Email: "ntnguyen@ubisen.com", Name: "Nguyen The Nguyen"}
			response := do_request("POST", url, user)
			body := parse_response(response)
			var responseData models.User
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)
			So(responseData.Email, ShouldEqual, user.Email)
			So(responseData.Name, ShouldEqual, user.Name)

			Convey("Create exist user should response status 400 and exist message", func() {
				response := do_request("POST", url, user)
				body := parse_response(response)
				var responseData Error
				err := json.Unmarshal(body, &responseData)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 400)
				So(responseData.Message, ShouldEqual, "This user has been exist!")
			})
		})
		Convey("Create with invalid email should return status 400 and email invalid message", func() {
			user := &models.User{Email: "invalidemail", Name: "invalidEmail"}
			response := do_request("POST", url, user)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Message, ShouldEqual, "Email invalid")
		})
		Convey("Create with empty email should return status 400 and email required message", func() {
			user := &models.User{Name: "Empty Email"}
			response := do_request("POST", url, user)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Message, ShouldEqual, "Email is required")
		})
	})
}
