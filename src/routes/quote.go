package routes

import (
	"net/http"

	"main/src/services"
	"main/src/routes/helper"

	"github.com/gin-gonic/gin"
)

// GetQuote gets a quote from the database
func GetQuote(c *gin.Context) {
	id, err := helper.ParamInt(c, "id")
	if helper.JSONError(c, err, "Quote ID must be an integer") { return }

	quote, err := services.GetQuote(id)
	if helper.JSONError(c, err, "Invalid quote ID") { return }

	c.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})
}

// GetRandomQuote gets a random quote from the database
func GetRandomQuote(c *gin.Context) {
	quote, err := services.GetRandomQuote()
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})
}

// GetQuotes gets a list of all the quotes from the database
func GetQuotes(c *gin.Context) {
	quotes, err := services.GetQuotes()
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"quotes": quotes,
	})
}
