package handlers

import (
	"github.com/golang-vietnam/forum/resources"
)

const (
	ITEMS_PER_PAGE = 15
)

var (
	userResource     = resources.NewResourceUser()
	categoryResource = resources.NewResourceCategory()
	postResource     = resources.NewResourcePost()
	authResource     = resources.NewResourceAuth()
)
