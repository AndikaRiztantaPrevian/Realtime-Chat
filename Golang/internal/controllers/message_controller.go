package controllers

import (
	"fmt"
	"net/http"

	"github.com/AndikaRiztantaPrevian/ChatApp/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/pusher/pusher-http-go/v5"
	"gorm.io/gorm"
)

var pusherClient = pusher.Client{
	AppID: "1819331",
	Key: "3bf5567f779653cead1b",
	Secret: "a0879c1c8f403abc2d18",
	Cluster: "ap1",
	Secure: true,
}

type MessageController struct {
    DB *gorm.DB
}

func (ctrl *MessageController) GetMessages(c *gin.Context) {
	var messages []models.Message
	if err := ctrl.DB.Find(&messages).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, messages)
}

func (ctrl *MessageController) CreateMessage(c *gin.Context) {
	var input struct {
		Message string `json:"message"`
	}
	
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message cannot be empty"})
		return
	}

	message := models.Message{Message: input.Message}

	if err := ctrl.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Created message with ID:", message.ID)

	if err := pusherClient.Trigger("chat-channel", "new-message", message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}
