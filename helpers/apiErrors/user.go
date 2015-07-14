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
	UserNotFound         = "USER_NOT_FOUND"
	UserIdInValid        = "USER_ID_INVALID"
	UserIdParamRequired  = "USER_ID_PARAM_REQUIRED"
)

var (
	userErrors = []apiError{
		apiError{
			Id:      UserIdParamRequired,
			Message: "userId in parameter required",
			Status:  400,
		},
		apiError{
			Id:      UserIdInValid,
			Message: "userId must objectId",
			Status:  400,
		},
		apiError{
			Id:      UserNotFound,
			Message: "This user not found",
			Status:  404,
		},
		apiError{
			Id:      UserExist,
			Message: "This user has been exist!",
			Status:  400,
		},
		apiError{
			Id:      UserEmailRequied,
			Message: "Email is required",
			Status:  400,
		},
		apiError{
			Id:      UserEmailInvalid,
			Message: "Email invalid",
			Status:  400,
		},
		apiError{
			Id:      UserEmailMin,
			Message: "Email min length is 3",
			Status:  400,
		},
		apiError{
			Id:      UserEmailMax,
			Message: "Email max length is 50",
			Status:  400,
		},
		apiError{
			Id:      UserNotLogined,
			Message: "You must login to do this",
			Status:  401,
		},
		apiError{
			Id:      UserRoleMin,
			Message: "Role min is 0 as public user",
			Status:  400,
		},
		apiError{
			Id:      UserRoleMax,
			Message: "Role max is 2 as admin",
			Status:  400,
		},
		apiError{
			Id:      UserPasswordRequired,
			Message: "Password is required",
			Status:  400,
		},
	}
)
