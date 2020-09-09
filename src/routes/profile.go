package routes

import (
	"net/http"

	app "main/src"
	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// GetProfileInfo gets a user's profile information
func GetProfileInfo(c *gin.Context) {
	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	email, joinTimestamp, err := services.GetProfileInfo(sessionID)
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"info": app.Object{
			"email": email,
			"joinTimestamp": joinTimestamp,
		},
	})
}

// ChangePassword changes the user's password
func ChangePassword(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "oldPassword", "newPassword")
	if failure { return }

	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	err := services.ChangePassword(sessionID, params["oldPassword"], params["newPassword"])
	if helper.JSONErrorDefault(c, err) { return }

	helper.JSONSuccess(c)
}
