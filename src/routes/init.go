package routes

import (
	"fmt"
	app "main/src"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoadRoutes creates all available routes
func LoadRoutes(router *gin.Engine, path string) {
	api := router.Group(path)

	api.GET("/person/id/:id", GetPerson)
	api.GET("/person/random", GetRandomPerson)
	api.GET("/people",        GetPeople)

	api.GET("/quote/id/:id",  GetQuote)
	api.GET("/quote/random",  GetRandomQuote)
	api.GET("/quotes",        GetQuotes)
}

// LoadErrorRoutes creates redirects to error pages when errors occur
func LoadErrorRoutes(router *gin.Engine, path string) {
	errorHandlers := map[int]gin.HandlerFunc{
		http.StatusNotFound: NotFound,
	}

	errors := router.Group(path)

	for statusCode, errorHandler := range errorHandlers {
		errors.GET(fmt.Sprintf("/%v", statusCode), errorHandler)
	}

	router.NoRoute(app.ReverseProxy(fmt.Sprintf("%s/%v", path, http.StatusNotFound)))
}
