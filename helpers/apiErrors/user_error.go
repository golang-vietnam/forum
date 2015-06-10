package apiErrors

var (
	USER_EXIST = Error{
		Id:      "USER_EXIST",
		Message: "This user has been exist!",
		Status:  400,
	}
	USER_EMAIL_REQUIRED = Error{
		Id:      "USER_EMAIL_REQUIRED",
		Message: "Email is required",
		Status:  400,
	}
	USER_EMAIL_INVALID = Error{
		Id:      "USER_EMAIL_INVALID",
		Message: "Email invalid",
		Status:  400,
	}
	USER_EMAIL_MIN = Error{
		Id:      "USER_EMAIL_MIN",
		Message: "Email min length is 3",
		Status:  400,
	}
	USER_EMAIL_MAX = Error{
		Id:      "USER_EMAIL_MAX",
		Message: "Email max length is 50",
		Status:  400,
	}
	USER_NOT_LOGINED = Error{
		Id:      "USER_NOT_LOGINED",
		Message: "You must login to do this",
		Status:  401,
	}
	USER_ROLE_MIN = Error{
		Id:      "USER_ROLE_MIN",
		Message: "Role min is 0 as public user",
		Status:  400,
	}
	USER_ROLE_MAX = Error{
		Id:      "USER_ROLE_MAX",
		Message: "Role max is 2 as editor",
		Status:  400,
	}

	USER_PASSWORD_REQUIRED = Error{
		Id:      "USER_PASSWORD_REQUIRED",
		Message: "Password is required",
		Status:  400,
	}
)
