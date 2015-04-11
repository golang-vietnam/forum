package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
}

func (u *User) Index(c *gin.Context) {
	c.String(http.StatusOK, "User page")
}
