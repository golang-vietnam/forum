package apiErrors

type Error struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) WithNewMessage(message string) *Error {
	return &Error{
		Id:      e.Id,
		Message: message,
		Status:  e.Status,
	}
}

var (
	SERVER_ERROR = Error{
		Id:      "SERVER_ERROR",
		Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
		Status:  500,
	}
	SERVER_TEMPORARILY_UNAVAILBLE = Error{
		Id:      "SERVER_TEMPORARILY_UNAVAILBLE",
		Message: "The server is temporarily unavailable.",
		Status:  500,
	}
)
