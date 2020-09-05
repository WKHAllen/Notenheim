package helper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// JSONSuccess sends a success message in JSON format
func JSONSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

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

// JSONUnexpectedError sends a generic error message back and prints a more helpful message
func JSONUnexpectedError(c *gin.Context, err error, errorMsg string) bool {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "An unexpected error occurred",
		})
		fmt.Printf("Unexpected error: %v\n", err)
	}
	return err != nil
}

// ParamInt gets a URL parameter as an integer
func ParamInt(c *gin.Context, paramName string) (int, error) {
	return strconv.Atoi(c.Param(paramName))
}

// QueryJSONError sends an error message in JSON format if a query parameter is missing
func QueryJSONError(c *gin.Context, key string) (string, bool) {
	value := c.Query(key)
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("Missing query parameter: %s", key),
		})
	}
	return value, value == ""
}

// QueriesJSONError attempts to get a sequence of query parameters and sends an error message if any are missing
func QueriesJSONError(c *gin.Context, keys ...string) (map[string]string, bool) {
	params := make(map[string]string)
	missing := []string{}
	
	for _, key := range keys {
		if c.Query(key) != "" {
			params[key] = c.Query(key)
		} else {
			missing = append(missing, key)
		}
	}

	switch (len(missing)) {
		case 0:
			return params, false
		case 1:
			c.JSON(http.StatusOK, gin.H{
				"error": fmt.Sprintf("Missing query parameter: %s", missing[0]),
			})
			return nil, true
		default:
			c.JSON(http.StatusOK, gin.H{
				"error": fmt.Sprintf("Missing query parameters: %s", strings.Join(missing, ", ")),
			})
			return nil, true
	}
}
