package resources

import (
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/mgo.v2/bson"
)

type resourceCategoryInterface interface {
	GetAll() ([]models.Category, error)
	GetById(id bson.ObjectId) (models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Remove(id bson.ObjectId) error
}

func NewResourceCategory() resourceCategoryInterface {
	return &resourceCategory{}
}

type resourceCategory struct {
}

const categoryColName = models.CategoryColName

func (r *resourceCategory) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := collection(categoryColName).Find(nil).All(&categories)
	return categories, err
}

func (r *resourceCategory) GetById(id bson.ObjectId) (models.Category, error) {
	var category models.Category
	err := collection(categoryColName).FindId(id).One(&category)
	return category, err
}

func (r *resourceCategory) Create(category *models.Category) error {
	category.Id = bson.NewObjectId()
	err := collection(categoryColName).Insert(category)
	return err
}

func (r *resourceCategory) Update(category *models.Category) error {
	err := collection(categoryColName).Update(bson.M{"_id": category.Id}, category)
	return err
}

func (r *resourceCategory) Remove(id bson.ObjectId) error {
	err := collection(categoryColName).RemoveId(id)
	return err
}
