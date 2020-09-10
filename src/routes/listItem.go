package routes

import (
	"net/http"
	"strconv"

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

// EditListItem edits a list item's content
func EditListItem(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "listItemID", "newContent")
	if failure { return }

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	err := services.EditListItem(sessionID, params["listItemID"], params["newContent"])
	if helper.JSONErrorDefault(c, err) { return }

	helper.JSONSuccess(c)
}

// CheckListItem checks or unchecks a list item
func CheckListItem(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "listItemID", "checked")
	if failure { return }

	var checked bool
	switch (params["checked"]) {
		case "true":
			checked = true
		case "false":
			checked = false
		default:
			c.JSON(http.StatusOK, gin.H{
				"error": "Parameter 'checked' must be 'true' or 'false'",
			})
			return
	}

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	err := services.CheckListItem(sessionID, params["listItemID"], checked)
	if helper.JSONErrorDefault(c, err) { return }

	helper.JSONSuccess(c)
}

// MoveListItem moves a list item up or down
func MoveListItem(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "listItemID", "newPosition")
	if failure { return }

	newPosition, err := strconv.Atoi(params["newPosition"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Parameter 'newPosition' must be an integer",
		})
		return
	}

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	err = services.MoveListItem(sessionID, params["listItemID"], newPosition)
	if helper.JSONErrorDefault(c, err) { return }

	helper.JSONSuccess(c)
}

// DeleteListItem deletes a list item
func DeleteListItem(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "listItemID")
	if failure { return }

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	err := services.DeleteListItem(sessionID, params["listItemID"])
	if helper.JSONErrorDefault(c, err) { return }

	helper.JSONSuccess(c)
}
