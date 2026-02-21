package controller

import (
	application "chat/Src/Endpoint/Message/Application"
	entities "chat/Src/Endpoint/Message/Domain/Entities"

	"github.com/gin-gonic/gin"
)

type SendMessageController struct {
	usecase application.SendMessageUseCase
}

func NewSendMessageController(usecase application.SendMessageUseCase) *SendMessageController {
	return &SendMessageController{
		usecase: usecase,
	}
}

func(c *SendMessageController) SendMessageController(context *gin.Context){
	var message *entities.Message

	if err := context.ShouldBindJSON(&message); err != nil {
		context.JSON(400, gin.H{"error": "Error 404 - Solicitud incorrecta, el cuerpo de la solicitud no es válido", "Detail": err.Error()})
		return
	}

	err := c.usecase.Execute(message)
	if err != nil {
		context.JSON(500, gin.H{"error": "Error 500 - Error interno del servidor", "Detail": err.Error()})
		return
	}

	var responseData = gin.H{
		"idMessage": message.ID,
		"senderId": message.SenderId,
		"receiveId": message.ReceiveId,
		"content": message.Content,
		"timeMessage": message.TimeMessage,
	}
	context.JSON(200, gin.H{"data": responseData})
}