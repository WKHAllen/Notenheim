package routes

import (
	"net/http"
	"os"

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

	verifyID, err := services.CreateVerification(params["email"])
	if helper.JSONErrorDefault(c, err) { return }

	services.PruneVerification(verifyID)

	helper.JSONSuccess(c)
}

// Login logs a user in
func Login(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "email", "password")
	if failure { return }

	sessionID, err := services.Login(params["email"], params["password"])
	if helper.JSONErrorDefault(c, err) { return }

	var domain string = "localhost"
	if os.Getenv("DEBUG") == "false" {
		domain = "notenheim.com"
	}

	c.SetCookie("sessionID", sessionID, 0, "/", domain, false, true)
	c.SetCookie("loggedIn", "true", 0, "/", domain, false, false)

	helper.JSONSuccess(c)
}

// Logout logs a user out
func Logout(c *gin.Context) {
	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Not logged in",
		})
	} else {
		var domain string = "localhost"
		if os.Getenv("DEBUG") == "false" {
			domain = "notenheim.com"
		}

		c.SetCookie("sessionID", "", -1, "/", domain, false, true)
		c.SetCookie("loggedIn", "false", 0, "/", domain, false, false)
		err := services.Logout(sessionID)
		if helper.JSONErrorDefault(c, err) { return }

		helper.JSONSuccess(c)
	}
}
