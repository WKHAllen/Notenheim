package routes

import (
	"os"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// Cookie maximum age
const cookieAge = 33 * 365.242 * 24 * 60 * 60 * 1000

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
	sessionID, err := c.Cookie("sessionID")
	if err == nil {
		err = services.Logout(sessionID)
		if helper.JSONErrorDefault(c, err) { return }
	}

	params, failure := helper.QueriesJSONError(c, "email", "password")
	if failure { return }

	sessionID, err = services.Login(params["email"], params["password"])
	if helper.JSONErrorDefault(c, err) { return }

	var domain string = "localhost"
	if os.Getenv("DEBUG") == "false" {
		domain = "notenheim.com"
	}

	c.SetCookie("sessionID", sessionID, cookieAge, "/", domain, false, false)

	helper.JSONSuccess(c)
}

// Logout logs a user out
func Logout(c *gin.Context) {
	sessionID, failure := helper.GetSessionID(c)
	if failure { return }

	var domain string = "localhost"
	if os.Getenv("DEBUG") == "false" {
		domain = "notenheim.com"
	}

	c.SetCookie("sessionID", "", -1, "/", domain, false, false)
	err := services.Logout(sessionID)
	if helper.JSONErrorDefault(c, err) { return }

	helper.JSONSuccess(c)
}
