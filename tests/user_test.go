package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/helpers"
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

		url := server + "/user/"
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
				var responseData helpers.Errors
				err := json.Unmarshal(body, &responseData)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 400)
				So(responseData.Errors[0].Title, ShouldEqual, "Bad Request")
				So(responseData.Errors[0].Detail, ShouldEqual, "This user has been exist!")
			})
		})
	})
}
