package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Gorm(gormCtx string, db gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(gormCtx, db)
		c.Next()
		db.Close()
	}
}
