package tests

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	categoryValidData = &categoryModel{
		Name:        "category",
		Description: "description",
		Slug:        "slug",
		Image:       "image_url",
	}
)

func CloneCategoryModel(category *categoryModel) *categoryModel {
	newCategory := *category
	return &newCategory
}

func TestCategoryApi(t *testing.T) {
	Convey("Test category api", t, func() {
		Reset(func() {
			clearAll()
		})

		Convey("Create category should successful", func() {
			category := CloneCategoryModel(categoryValidData)
			response := do_request("POST", categoryApi, category)
			body := parse_response(response)
			var responseCategory categoryModel
			err := json.Unmarshal(body, &responseCategory)
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 201)
			So(responseCategory.Description, ShouldEqual, category.Description)
			So(responseCategory.Name, ShouldEqual, category.Name)
			So(responseCategory.Slug, ShouldEqual, category.Slug)
			So(responseCategory.Image, ShouldEqual, category.Image)

		})
	})
}
