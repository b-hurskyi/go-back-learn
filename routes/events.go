package routes

import (
	"net/http"
	"strconv"

	"github.com/b-hurskyi/go-back-learn/models"
	"github.com/b-hurskyi/go-back-learn/utils"
	"github.com/gin-gonic/gin"
)

func deleteEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(ctx.Request.Context(), eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}

func updateEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	_, err = models.GetEventById(ctx.Request.Context(), eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func getEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(ctx.Request.Context(), eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something wrong"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var event models.Event
	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserID = userId
	err = event.Save(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
