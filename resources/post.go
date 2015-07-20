package resources

import (
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type resourcePostInterface interface {
	ListAll() ([]models.Post, error)
	ListPaging(page int, perPage int) ([]models.Post, error)
	ListPagingByCategory(id bson.ObjectId, page int, perPage int) ([]models.Post, error)
	GetById(id bson.ObjectId) (models.Post, error)
	Create(u *models.Post) error
	Update(u *models.Post) error
	RemoveById(id bson.ObjectId) error
	ParseError(err error) []error
}

func NewResourcePost() resourcePostInterface {
	return &resourcePost{}
}

type resourcePost struct {
}

const postColName = models.PostColName

func (r *resourcePost) ListAll() ([]models.Post, error) {
	var posts []models.Post
	if err := collection(postColName).Find(nil).All(&posts); err != nil {
		if err == mgo.ErrNotFound {
			return nil, apiErrors.ThrowError(apiErrors.PostNotFound)
		}
		panic(err)
	}
	return posts, nil
}

//Please make sure page >= 0 && perPage >= 0
func (r *resourcePost) ListPaging(page int, perPage int) ([]models.Post, error) {
	var posts []models.Post
	if err := collection(postColName).Find(nil).Limit(perPage).Skip(perPage * page).All(&posts); err != nil {
		panic(err)
	}
	return posts, nil
}

//Please make sure page >= 0 && perPage >= 0
func (r *resourcePost) ListPagingByCategory(categoryId bson.ObjectId, page int, perPage int) ([]models.Post, error) {
	var posts []models.Post
	if err := collection(postColName).Find(bson.M{"category._id": categoryId}).Limit(perPage).Skip(perPage * page).All(&posts); err != nil {
		panic(err)
	}
	return posts, nil
}

func (r *resourcePost) GetById(id bson.ObjectId) (models.Post, error) {
	var post models.Post
	if err := collection(postColName).FindId(id).One(&post); err != nil {
		if err == mgo.ErrNotFound {
			return post, apiErrors.ThrowError(apiErrors.PostNotFound)
		}
		panic(err)
	}
	return post, nil
}

func (r *resourcePost) Create(p *models.Post) error {
	p.Id = bson.NewObjectId()
	if err := collection(postColName).Insert(p); err != nil {
		if mgo.IsDup(err) {
			return apiErrors.ThrowError(apiErrors.PostExist)
		}
		panic(err)
	}
	return nil
}

func (r *resourcePost) Update(p *models.Post) error {
	p.Id = bson.NewObjectId()
	if err := collection(postColName).UpdateId(p.Id, p); err != nil {
		if mgo.IsDup(err) {
			return apiErrors.ThrowError(apiErrors.PostExist)
		}
		panic(err)
	}
	return nil
}

func (r *resourcePost) RemoveById(id bson.ObjectId) error {
	if err := collection(postColName).RemoveId(id); err != nil {
		if err == mgo.ErrNotFound {
			return apiErrors.ThrowError(apiErrors.PostNotFound)
		}
		panic(err)
	}
	return nil
}

func (r *resourcePost) ParseError(err error) []error {
	var errors []error
	if errs, ok := err.(*validator.StructErrors); ok {
		for _, v := range errs.Errors {
			switch v.Field {
			case "User":
				switch v.Tag {
				case "required":
					errors = append(errors, apiErrors.ThrowError(apiErrors.PostUserRequired))
				}
			case "Title":
				switch v.Tag {
				case "required":
					errors = append(errors, apiErrors.ThrowError(apiErrors.PostTitleRequired))
				}
			case "Content":
				switch v.Tag {
				case "required":
					errors = append(errors, apiErrors.ThrowError(apiErrors.PostContentRequired))
				}
			}
		}
	} else {
		panic("Can not parse error")
	}

	return errors
}
