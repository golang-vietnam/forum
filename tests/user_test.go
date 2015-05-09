package tests

import (
	"github.com/golang-vietnam/forum/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUserNotLogin(t *testing.T) {
	server := getServer()
	t.Parallel()
	Convey("POST register account", t, func() {
		Convey("With invalid email should resond status 400 and correct error data", func() {
			user := &models.User{Email: "inValidEmail", Name: "User Name"}
			response := do_request("POST", server, user)
			So(response.StatusCode, ShouldEqual, 400)
		})
	})
}
