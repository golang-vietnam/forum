package middleware

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-vietnam/forum/config"
	"github.com/jinzhu/gorm"
)

func Gorm() gin.HandlerFunc {

	sqlConnection := config.GetDB("user") + ":" + config.GetDB("password") +
		"@tcp(" + config.GetDB("host") + ":" + config.GetDB("port") + ")/" + config.GetDB("name") + "?parseTime=True"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
		return nil
	}
	db.SingularTable(true)
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
		db.Close()
	}
}
