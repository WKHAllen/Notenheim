package routes

import (
	"net/http"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// Verify attempts to mark a user as verified
func Verify(c *gin.Context) {
	params, failure := helper.QueriesJSONError(c, "verifyID")
	if failure { return }

	err := services.Verify(params["verifyID"])
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}
