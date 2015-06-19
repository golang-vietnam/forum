package apiErrors

const (
	ApiErrorNotFound  = "API_ERROR_NOT_FOUND"
	ApiErrorIdRequied = "API_ERROR_ID_REQUIRED"
)

var (
	apiErrorErrors = []apiError{
		apiError{
			Id:      ApiErrorNotFound,
			Message: "API error not found",
			Status:  401,
		},
		apiError{
			Id:      ApiErrorIdRequied,
			Message: "API error Id required ",
		},
	}
)
