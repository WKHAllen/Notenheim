package routes

import (
	"net/http"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// NewList creates a new list
func NewList(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "title")
	if failure { return }

	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Not logged in",
		})
		return
	}

	listID, err := services.NewList(sessionID, params["title"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"listID": listID,
	})
}
