package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"statusCode": code, "data": data})
}

func MakeFail(c *gin.Context, code int, message interface{}) {
	c.JSON(http.StatusOK, gin.H{"statusCode": code, "message": message})
}
