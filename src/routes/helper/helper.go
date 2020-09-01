package helper

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JSONError sends an error message in JSON format if an error has occurred
func JSONError(c *gin.Context, err error, errorMsg string) bool {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": errorMsg,
		})
	}
	return err != nil
}

// JSONErrorDefault sends the given error message in JSON format if an error has occurred
func JSONErrorDefault(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	return err != nil
}

// ParamInt gets a URL parameter as an integer
func ParamInt(c *gin.Context, paramName string) (int, error) {
	return strconv.Atoi(c.Param(paramName))
}
