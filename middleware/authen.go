package middleware

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
	"github.com/golang-vietnam/forum/helpers/config"
	"github.com/golang-vietnam/forum/helpers/utils"
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

func (a *authMiddleware) RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt_lib.ParseFromRequest(c.Request, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(config.GetSecret()))
			return b, nil
		})

		if err != nil || token == nil || (token != nil && !token.Valid) {
			c.Error(apiErrors.ThrowError(apiErrors.AccessDenied))
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

func (a *authMiddleware) UserRequirePermission(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := utils.MustGetCurrentUser(c)
		if currentUser.Role < role {
			c.Error(apiErrors.ThrowError(apiErrors.AccessDenied))
			c.Abort()
			return
		}
		c.Next()
	}
}

func (a *authMiddleware) UserHasAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		currentUser := utils.MustGetCurrentUser(c)
		userData := utils.MustGetUserData(c)

		if currentUser.Role == models.NormalUser {
			if currentUser.Id != userData.Id {
				c.Error(apiErrors.ThrowError(apiErrors.AccessDenied))
				c.Abort()
				return
			}
		} else {
			if currentUser.Role <= userData.Role {
				c.Error(apiErrors.ThrowError(apiErrors.AccessDenied))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
