package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFound handles 404 codes
func NotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": "not found",
	})
}
