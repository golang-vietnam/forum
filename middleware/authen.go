package middleware

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
)

type authMiddlewareInterface interface {
	RequireLogin(secret string) gin.HandlerFunc
	RequirePermission(role int) gin.HandlerFunc
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
func (a *authMiddleware) RequireLogin(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := jwt_lib.ParseFromRequest(c.Request, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.Error(apiErrors.ThrowError(apiErrors.UserNotLogined))
			return
		}
		c.Set("user", user)
	}
}

/**

	TODO:
	- If user not login as gin context get "user" is nill -> user role = normalUser
	- If user has role < role param return access deny error
	- If user has role >= role param -> pass

**/
func (a *authMiddleware) RequirePermission(role int) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
