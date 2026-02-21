package router

import (
	controller "chat/Src/Endpoint/Message/Infrestructure/Controller"

	"github.com/gin-gonic/gin"
)

func MessageRouter(gin *gin.Engine,
	GetAllMessage *controller.GetAllMessageController,
	SendMessage *controller.SendMessageController,
	DeleteMessage *controller.DeleteMessageController) {

	messageGroup := gin.Group("api/v1/message")
	{
		messageGroup.GET("/getAll", GetAllMessage.GetAllMessagesController)
		messageGroup.POST("/send", SendMessage.SendMessageController)
		messageGroup.DELETE("/delete/:messageId", DeleteMessage.DeleteMessageController)
	}
	
}