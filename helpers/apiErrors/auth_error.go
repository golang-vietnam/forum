package apiErrors

var (
	AUTH_EMAIL_INVALID = Error{
		Id:      "AUTH_EMAIL_INVALID",
		Message: "Email login invalid",
		Status:  401,
	}
	AUTH_PASSWORD_INVALID = Error{
		Id:      "AUTH_PASSWORD_INVALID",
		Message: "Password login invalid",
		Status:  401,
	}
)
