package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-vietnam/forum/helpers/apiErrors"
)

type errorControllerInterface interface {
	Error404(c *gin.Context)
	List(c *gin.Context)
	GetById(c *gin.Context)
}
type errorController struct {
}

func NewErrorController() errorControllerInterface {
	return &errorController{}
}
func (e *errorController) Error404(c *gin.Context) {
	c.String(404, "Page not found")
}
func (e *errorController) List(c *gin.Context) {
	c.JSON(200, gin.H{"errors": apiErrors.ApiErrors})
}
func (e *errorController) GetById(c *gin.Context) {
	var errorId string
	if errorId = c.ParamValue("errorId"); errorId == "" {
		if errorId = c.Request.URL.Query().Get("errorId"); errorId == "" {
			c.Error(apiErrors.ThrowError(apiErrors.ApiErrorIdRequied))
			return
		}
	}
	apiError := apiErrors.FindErrorById(errorId)
	if apiError == nil {
		c.Error(apiErrors.ThrowError(apiErrors.ApiErrorNotFound))
		return
	}
	c.JSON(200, apiError)
}
