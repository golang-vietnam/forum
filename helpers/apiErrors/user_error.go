package apiErrors

const (
	UserExist            = "USER_EXIST"
	UserEmailRequied     = "USER_EMAIL_REQUIRED"
	UserEmailInvalid     = "USER_EMAIL_INVALID"
	UserEmailMin         = "USER_EMAIL_MIN"
	UserEmailMax         = "USER_EMAIL_MAX"
	UserNotLogined       = "USER_NOT_LOGINED"
	UserRoleMin          = "USER_ROLE_MIN"
	UserRoleMax          = "USER_ROLE_MAX"
	UserPasswordRequired = "USER_PASSWORD_REQUIRED"
)

var (
	userErrors = []*Error{
		&Error{
			Id:      UserExist,
			Message: "This user has been exist!",
			Status:  400,
		},
		&Error{
			Id:      UserEmailRequied,
			Message: "Email is required",
			Status:  400,
		},
		&Error{
			Id:      UserEmailInvalid,
			Message: "Email invalid",
			Status:  400,
		},
		&Error{
			Id:      UserEmailMin,
			Message: "Email min length is 3",
			Status:  400,
		},
		&Error{
			Id:      UserEmailMax,
			Message: "Email max length is 50",
			Status:  400,
		},
		&Error{
			Id:      UserNotLogined,
			Message: "You must login to do this",
			Status:  401,
		},
		&Error{
			Id:      UserRoleMin,
			Message: "Role min is 0 as public user",
			Status:  400,
		},
		&Error{
			Id:      UserRoleMax,
			Message: "Role max is 1 as editor",
			Status:  400,
		},
		&Error{
			Id:      UserPasswordRequired,
			Message: "Password is required",
			Status:  400,
		},
	}
)
