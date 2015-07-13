package middleware

import (
	"github.com/golang-vietnam/forum/resources"
)

var (
	userResource     = resources.NewResourceUser()
	postResource     = resources.NewResourcePost()
	authResource     = resources.NewResourceAuth()
	categoryResource = resources.NewResourceCategory()
)
