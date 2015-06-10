package middleware

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
)

/**

	TODO:
	- Check user is login
	- If not return not login error
	- If logined set "user" in context

**/

func RequireLogin(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := jwt_lib.ParseFromRequest(c.Request, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.Error(&apiErrors.USER_NOT_LOGINED)
			return
		}
		c.Set("user", user)
	}
}

/**

	TODO:
	- If user not login as gin context get "user" is nill -> user role = 0
	- If user has role < role param return access deny error
	- If user has role >= role param -> pass

**/

func RequirePermission(role int) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
