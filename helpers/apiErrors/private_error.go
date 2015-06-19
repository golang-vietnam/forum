package apiErrors

const (
	ServerError                 = "SERVER_ERROR"
	ServerTemporarilyUnavailble = "SERVER_TEMPORARILY_UNAVAILBLE"
)

var privateErrors = []*Error{
	&Error{
		Id:      ServerError,
		Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
		Status:  500,
	},
	&Error{
		Id:      ServerTemporarilyUnavailble,
		Message: "The server is temporarily unavailable.",
		Status:  500,
	},
}
