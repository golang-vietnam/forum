package apiErrors

const (
	CategoryNotFound     = "CATEGORY_NOT_FOUND"
	CategoryNameRequired = "CATEGORY_NAME_REQUIRED"
	CategorySlugRequired = "CATEGORY_SLUG_REQUIRED"
	CategoryExist        = "CATEGORY_EXIST"
)

var categoryErrors = []apiError{
	apiError{
		Id:      CategoryNotFound,
		Message: "Category not found",
		Status:  404,
	},
	apiError{
		Id:      CategorySlugRequired,
		Message: "Category slug required",
		Status:  400,
	},
	apiError{
		Id:      CategoryNameRequired,
		Message: "Category name required",
		Status:  400,
	},
	apiError{
		Id:      CategoryExist,
		Message: "Category exist",
		Status:  400,
	},
}
