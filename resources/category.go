package resources

import (
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type resourceCategoryInterface interface {
	GetAll() []*models.Category
	GetById(id string) (*models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Remove(id bson.ObjectId) error
	ParseError(err error) []error
}

func NewResourceCategory() resourceCategoryInterface {
	return &resourceCategory{}
}

type resourceCategory struct {
}

const categoryColName = models.CategoryColName

func (r *resourceCategory) GetAll() []*models.Category {
	var categories []*models.Category
	if err := collection(categoryColName).Find(nil).All(&categories); err != nil {
		panic(err)
	}
	return categories
}

func (r *resourceCategory) GetById(id string) (*models.Category, error) {

	if !bson.IsObjectIdHex(id) {
		return nil, apiErrors.ThrowError(apiErrors.IdInvalid)
	}

	var category models.Category
	if err := collection(categoryColName).FindId(bson.ObjectIdHex(id)).One(&category); err != nil {
		if err == mgo.ErrNotFound {
			return nil, apiErrors.ThrowError(apiErrors.CategoryNotFound)
		}
		panic(err)
	}
	return &category, nil
}

func (r *resourceCategory) Create(category *models.Category) error {
	category.Id = bson.NewObjectId()
	if err := collection(categoryColName).Insert(category); err != nil {
		if mgo.IsDup(err) {
			return apiErrors.ThrowError(apiErrors.CategoryExist)
		}
		panic(err)
	}
	return nil
}

func (r *resourceCategory) Update(category *models.Category) error {
	err := collection(categoryColName).Update(bson.M{"_id": category.Id}, category)
	return err
}

func (r *resourceCategory) Remove(id bson.ObjectId) error {
	err := collection(categoryColName).RemoveId(id)
	return err
}
func (r *resourceCategory) ParseError(err error) []error {
	var errors []error
	if errs, ok := err.(*validator.StructErrors); ok {
		for _, v := range errs.Errors {
			switch v.Field {
			case "Name":
				switch v.Tag {
				case "required":
					errors = append(errors, apiErrors.ThrowError(apiErrors.CategoryNameRequired))
				}
			case "Slug":
				switch v.Tag {
				case "required":
					errors = append(errors, apiErrors.ThrowError(apiErrors.CategorySlugRequired))
				}
			}
		}
	} else {
		panic("Can not parse error")
	}

	return errors
}
