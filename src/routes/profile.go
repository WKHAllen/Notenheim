package routes

import (
	"net/http"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// ChangePassword changes the user's password
func ChangePassword(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "newPassword")
	if failure { return }

	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Not logged in",
		})
		return
	}

	err = services.ChangePassword(sessionID, params["newPassword"])
	if helper.JSONErrorDefault(c, err) { return }

	helper.JSONSuccess(c)
}
