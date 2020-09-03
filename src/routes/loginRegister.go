package routes

import (
	"net/http"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// Register registers a user
func Register(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "email", "password")
	if failure { return }

	err := services.Register(params["email"], params["password"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

// Login logs a user in
func Login(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "email", "password")
	if failure { return }

	sessionID, err := services.Login(params["email"], params["password"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error":     nil,
		"sessionID": sessionID,
	})
}

// Logout logs a user out
func Logout(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "sessionID")
	if failure { return }

	err := services.Logout(params["sessionID"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}
