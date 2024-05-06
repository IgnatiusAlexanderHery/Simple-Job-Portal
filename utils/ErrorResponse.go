package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessfulResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func BadRequestError(c *gin.Context, text string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": text})
}

func UnauthorizedError(c *gin.Context, text string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Unauthorized": text})
}

func NotFoundError(c *gin.Context, text string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Not Found": text})
}
