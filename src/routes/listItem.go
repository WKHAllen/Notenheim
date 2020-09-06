package routes

import (
	"net/http"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// NewListItem creates a new item in a list
func NewListItem(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "listID")
	if failure { return }

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	listItemID, err := services.NewListItem(sessionID, params["listID"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"listItemID": listItemID,
	})
}
