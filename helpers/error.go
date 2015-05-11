package helpers

type Errors struct {
	Errors []Error `json:"errors"`
}

func (e *Errors) StatusCode() int {
	statusCode := 500
	if len(e.Errors) > 0 {
		statusCode = e.Errors[0].Status
	}
	return statusCode
}

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (e *Error) IsNil() bool {
	if e.Status == 0 && e.Detail == "" && e.Title == "" {
		return true
	}
	return false
}

var (
	ErrInternalServer = Error{500, "Internal Server Error", "Something went wrong."}
	ErrBadRequest     = Error{400, "Bad Request", "The request had bad syntax or was inherently impossible to be satisfied"}
	ErrUnauthorized   = Error{401, "Unauthorized", "Login required"}
	ErrForbidden      = Error{403, "Forbidden", "Access deny"}
	ErrNotFound       = Error{404, "Not Found", "Not found anything matching the URI given"}
)
