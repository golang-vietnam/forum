package middleware

import (
	// "fmt"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/config"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/models"
)

type authMiddlewareInterface interface {
	RequireLogin() gin.HandlerFunc
	UserRequirePermission(role int) gin.HandlerFunc
	UserHasAuthorization() gin.HandlerFunc
}
type authMiddleware struct {
}

func NewAuthMiddleware() authMiddlewareInterface {
	return &authMiddleware{}
}

/**

	TODO:
	- Check user is login
	- If not return not login error
	- If logined set "user" in context

**/
func (a *authMiddleware) RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt_lib.ParseFromRequest(c.Request, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(config.GetSecret()))
			return b, nil
		})

		if err != nil || token == nil || (token != nil && !token.Valid) {
			c.Error(apiErrors.ThrowError(apiErrors.UserNotLogined))
			c.Abort()
			return
		}

		var currentUser *models.User
		var findUserErr error

		if userId, ok := token.Claims["userId"].(string); ok {
			if currentUser, findUserErr = userResource.GetById(userId); findUserErr != nil {
				c.Error(findUserErr)
				c.Abort()
				return
			}
		} else {
			panic("Must load userId in token")
		}
		c.Set("currentUser", currentUser)
		c.Next()
	}
}

/**

	TODO:
	- If user has role < role param return access deny error
	- If user has role >= role param -> pass

**/
func (a *authMiddleware) UserRequirePermission(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.MustGet("currentUser").(*models.User)
		if !ok {
			panic("data with key currentUser must models.User type")
		}
		if user.Role < role {
			c.Error(apiErrors.ThrowError(apiErrors.AccessDenied))
			c.Abort()
			return
		}
		c.Next()
	}
}

func (a *authMiddleware) UserHasAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		currentUser, currentUserOk := c.MustGet("currentUser").(*models.User)
		if !currentUserOk {
			panic("data with key currentUser must models.User type")
		}

		userData, userDataOk := c.MustGet("userData").(*models.User)
		if !userDataOk {
			panic("data with key userData must models.User type")
		}

		if currentUser.Id != userData.Id {
			c.Error(apiErrors.ThrowError(apiErrors.AccessDenied))
			c.Abort()
			return
		}

		c.Next()
	}
}
