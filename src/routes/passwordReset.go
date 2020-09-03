package routes

import (
	"net/http"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// RequestPasswordReset requests a password reset
func RequestPasswordReset(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "email")
	if failure { return }

	services.RequestPasswordReset(params["email"])

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

// ValidPasswordResetID verifies that a password reset ID is valid
func ValidPasswordResetID(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "resetID")
	if failure { return }

	valid := services.ValidPasswordResetID(params["resetID"])
	if !valid {
		c.JSON(http.StatusOK, gin.H{
			"error": "Invalid password reset ID",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": nil,
		})
	}
}

// ResetPassword resets a user's password
func ResetPassword(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "resetID", "newPassword")
	if failure { return }

	err := services.ResetPassword(params["resetID"], params["newPassword"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}
