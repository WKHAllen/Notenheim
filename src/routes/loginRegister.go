package routes

import (
	"net/http"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// Register registers a user
func Register(c * gin.Context) {
	params, failure := helper.QueriesJSONError(c, "email", "password")
	if failure { return }

	err := services.Register(params["email"], params["password"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
