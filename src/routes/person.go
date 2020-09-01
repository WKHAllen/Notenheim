package routes

import (
	"net/http"

	"main/src/routes/helper"
	"main/src/services"

	"github.com/gin-gonic/gin"
)

// GetPerson gets a person from the database
func GetPerson(c *gin.Context) {
	id, err := helper.ParamInt(c, "id")
	if helper.JSONError(c, err, "Person ID must be an integer") { return }

	person, err := services.GetPerson(id)
	if helper.JSONError(c, err, "Invalid person ID") { return }

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

// GetRandomPerson gets a random person from the database
func GetRandomPerson(c *gin.Context) {
	person, err := services.GetRandomPerson()
	if helper.JSONErrorDefault(c, err) { return}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

// GetPeople gets a list of all the people from the database
func GetPeople(c *gin.Context) {
	people, err := services.GetPeople()
	if helper.JSONErrorDefault(c, err) { return }

	c.JSON(http.StatusOK, gin.H{
		"people": people,
	})
}
