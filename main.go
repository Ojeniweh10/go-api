package main

import (
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", postEvents)
	server.GET("/events/:id", getEventsByID)

	server.Run(":8080") //localhost:8080

}

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, models.Event)
}

func postEvents(c *gin.Context) {
	var event models.Events
	if err := c.BindJSON(&event); err != nil {
		return
	}
	models.Event = append(models.Event, event)
	c.IndentedJSON(http.StatusCreated, event)
}

func getEventsByID(c *gin.Context) {
	id := c.Param("id")
	iD, _ := strconv.Atoi(id)
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range models.Event {
		if a.ID == iD {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
