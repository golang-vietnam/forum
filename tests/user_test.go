package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/database"
	"github.com/golang-vietnam/forum/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	userColName = "user"
)

func TestUser(t *testing.T) {
	server := getServer()

	database.InitDb()
	defer database.CloseDb()
	database.ClearAllUser()

	Convey("POST create user", t, func() {
		url := server + "/v1/user/"
		Convey("Create not exist user should response status 201 and correct user data.", func() {

			user := &models.User{Email: "ntnguyen@ubisen.com", Name: "Nguyen The Nguyen", Password: "golang", Role: 2}
			response := do_request("POST", url, user)
			body := parse_response(response)
			var responseUser models.User
			err := json.Unmarshal(body, &responseUser)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)
			So(responseUser.Email, ShouldEqual, user.Email)
			So(responseUser.Name, ShouldEqual, user.Name)
			So(responseUser.Role, ShouldEqual, 0)

			var userInDb models.User
			database.Collection(userColName).FindId(responseUser.Id).One(&userInDb)
			So(userInDb.Id, ShouldEqual, responseUser.Id)

			Convey("Create exist user should response status 400 and exist message", func() {
				response := do_request("POST", url, user)
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
			user := &models.User{Email: "invalidemail", Name: "invalidEmail", Password: "invalidEmail"}
			response := do_request("POST", url, user)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Id, ShouldEqual, "USER_EMAIL_INVALID")
			So(responseData.Message, ShouldEqual, "Email invalid")
		})
		Convey("Create with empty email should return status 400 and email required message", func() {
			user := &models.User{Name: "Empty Email", Password: "Empty Email"}
			response := do_request("POST", url, user)
			body := parse_response(response)
			var responseData Error
			err := json.Unmarshal(body, &responseData)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 400)
			So(responseData.Id, ShouldEqual, "USER_EMAIL_REQUIRED")
			So(responseData.Message, ShouldEqual, "Email is required")
		})
		Convey("Create with empty password should return status 400 and password required message", func() {
			user := &models.User{Name: "Ntn", Email: "ntnguyen@ubisen.com"}
			response := do_request("POST", url, user)
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
