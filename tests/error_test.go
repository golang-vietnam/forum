package tests

import (
	"encoding/json"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

type ApiErrors struct {
	Errors []apiErrors.ApiError `json:"errors"`
}

func TestErrorApi(t *testing.T) {
	Convey("Get API Error", t, func() {

		Convey("Get all api error should return list api error and status 200", func() {
			response := do_request("GET", errorApi, nil)
			body := parse_response(response)
			var responseApiErrors ApiErrors
			err := json.Unmarshal(body, &responseApiErrors)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, http.StatusOK)
			So(len(responseApiErrors.Errors), ShouldEqual, len(apiErrors.ApiErrors))
		})

		Convey("Get an api error should return an api error and status 200", func() {
			response := do_request("GET", errorApi+apiErrors.ServerError, nil)
			body := parse_response(response)
			var responseApiError apiErrors.ApiError
			err := json.Unmarshal(body, &responseApiError)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, http.StatusOK)
			So(responseApiError.Id, ShouldEqual, apiErrors.ServerError)
		})

		Convey("Get invalid error Id return not api not found and status 404", func() {
			response := do_request("GET", errorApi+"NO_ERROR", nil)
			body := parse_response(response)
			var responseApiError apiErrors.ApiError
			err := json.Unmarshal(body, &responseApiError)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, http.StatusNotFound)
			So(responseApiError.Id, ShouldEqual, apiErrors.ApiErrorNotFound)
		})

	})
}
