package routes

import (
	"github.com/AndikaRiztantaPrevian/ChatApp/internal/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	messageController := controllers.MessageController{DB: db}

	router.GET("/api/getMessage", messageController.GetMessages)
	router.POST("/api/createMessage", messageController.CreateMessage)

	return router
}