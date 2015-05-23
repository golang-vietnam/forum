package helpers

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
)

type Errors struct {
	Errors []*Error `json:"errors"`
}

func (e *Errors) StatusCode() int {
	if len(e.Errors) > 0 {
		return e.Errors[0].Status
	}
	return 500
}

func (e *Errors) HasError() bool {
	return len(e.Errors) > 0
}

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (e *Error) IsNil() bool {
	return e.Status == 0 && e.Detail == "" && e.Title == ""
}

var (
	ErrInternalServer = Error{500, "Internal Server Error", "Something went wrong."}
	ErrBadRequest     = Error{400, "Bad Request", "The request had bad syntax or was inherently impossible to be satisfied"}
	ErrUnauthorized   = Error{401, "Unauthorized", "Login required"}
	ErrForbidden      = Error{403, "Forbidden", "Access deny"}
	ErrNotFound       = Error{404, "Not Found", "Not found anything matching the URI given"}
)

func SetErrors(c *gin.Context) {
	c.Set("errors", &Errors{})
}

func AddError(c *gin.Context, e Error) {
	errors := c.MustGet("errors").(*Errors)
	errors.Errors = append(errors.Errors, &e)
}
func GetErrors(c *gin.Context) *Errors {
	return c.MustGet("errors").(*Errors)
}
