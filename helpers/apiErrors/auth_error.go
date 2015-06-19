package apiErrors

const (
	AuthEmailInvalid    = "AUTH_EMAIL_INVALID"
	AuthPasswordInValid = "AUTH_PASSWORD_INVALID"
)

var (
	authErrors = []*Error{
		&Error{
			Id:      AuthEmailInvalid,
			Message: "Email login invalid",
			Status:  401,
		},
		&Error{
			Id:      AuthPasswordInValid,
			Message: "Password login invalid",
			Status:  401,
		},
	}
)
