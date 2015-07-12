package apiErrors

const (
	CategoryNotFound = "CATEGORY_NOT_FOUND"
)

var categoryErrors = []apiError{
	apiError{
		Id:      CategoryNotFound,
		Message: "Category not found",
		Status:  404,
	},
}
