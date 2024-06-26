package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example.com/rest-api/models"
)

func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func GetEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID format"})
		return
	}

	event, err := models.GetEventByID(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := c.Get("userId")
	event.UserID = userId.(int64)

	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func UpdateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID format"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this event"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})
}

func DeleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID format"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
