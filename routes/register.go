package routes

import (
	"net/http"
	"strconv"

	"github.com/b-hurskyi/go-back-learn/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(ctx.Request.Context(), eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(ctx.Request.Context(), eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registration canceled."})
}
