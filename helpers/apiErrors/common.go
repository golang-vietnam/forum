package apiErrors

const (
	ServerError                 = "SERVER_ERROR"
	ServerTemporarilyUnavailble = "SERVER_TEMPORARILY_UNAVAILBLE"
	AccessDenied                = "ACCESS_DENIED"
)

var privateErrors = []apiError{
	apiError{
		Id:      ServerError,
		Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
		Status:  500,
	},
	apiError{
		Id:      ServerTemporarilyUnavailble,
		Message: "The server is temporarily unavailable.",
		Status:  500,
	},
	apiError{
		Id:      AccessDenied,
		Message: "Access Denied",
		Status:  403,
	},
}
