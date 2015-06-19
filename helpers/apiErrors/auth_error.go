package apiErrors

const (
	AuthEmailInvalid    = "AUTH_EMAIL_INVALID"
	AuthPasswordInValid = "AUTH_PASSWORD_INVALID"
)

var (
	authErrors = []*apiError{
		&apiError{
			Id:      AuthEmailInvalid,
			Message: "Email login invalid",
			Status:  401,
		},
		&apiError{
			Id:      AuthPasswordInValid,
			Message: "Password login invalid",
			Status:  401,
		},
	}
)
