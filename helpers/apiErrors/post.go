package apiErrors

const (
	PostExist           = "POST_EXIST"
	PostNotFound        = "POST_NOT_FOUND"
	PostIdInValid       = "POST_ID_INVALID"
	PostIdParamRequired = "POST_ID_PARAM_REQUIRED"
)

var (
	postErrors = []apiError{
		apiError{
			Id:      PostIdParamRequired,
			Message: "postId in parameter required",
			Status:  400,
		},
		apiError{
			Id:      PostIdInValid,
			Message: "postId must objectId",
			Status:  400,
		},
		apiError{
			Id:      PostNotFound,
			Message: "This post not found",
			Status:  404,
		},
		apiError{
			Id:      PostExist,
			Message: "This post has been exist!",
			Status:  400,
		},
	}
)
