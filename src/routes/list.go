package routes

import (
	"net/http"

	app "main/src"
	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// NewList creates a new list
func NewList(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "title")
	if failure { return }

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	listID, err := services.NewList(sessionID, params["title"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"listID": listID,
	})
}

// GetLists gets all lists created by a user
func GetLists(c *gin.Context) {
	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	lists, err := services.GetLists(sessionID)
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"lists": lists,
	})
}

// ListInfo gets the title and items within a list
func ListInfo(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "listID")
	if failure { return }

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	title, items, err := services.ListInfo(sessionID, params["listID"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"info": app.Object{
			"title": title,
			"items": items,
		},
	})
}

// DeleteList deletes a list
func DeleteList(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "listID")
	if failure { return }

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	err := services.DeleteList(sessionID, params["listID"])
	if helper.JSONErrorDefault(c, err) { return }

	helper.JSONSuccess(c)
}
