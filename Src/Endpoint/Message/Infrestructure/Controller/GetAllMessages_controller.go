package controller

import (
	application "chat/Src/Endpoint/Message/Application"
	"strconv"
	"github.com/gin-gonic/gin"
)

type GetAllMessageController struct {
	usecase *application.GetMessageByUserIdUseCase
}

func NewGetAllMessageController(usecase *application.GetMessageByUserIdUseCase) *GetAllMessageController {
	return &GetAllMessageController{
		usecase: usecase,
	}
}

func (c *GetAllMessageController) GetAllMessagesController(ctx *gin.Context) {
    id := ctx.Query("userId")
    if id == "" {
        ctx.JSON(400, gin.H{"error": "userId es requerido"})
        return
    }

    idInt, err := strconv.Atoi(id)
    if err != nil {
        ctx.JSON(400, gin.H{"error": "userId debe ser un número entero"})
        return
    }

    messages, err := c.usecase.Execute(idInt)
    if err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }

    var responseData []gin.H
    for _, message := range messages {
        responseData = append(responseData, gin.H{
            "idMessage":   message.ID,
            "senderId":    message.SenderId,
            "receiveId":   message.ReceiveId,
            "content":     message.Content,
            "timeMessage": message.TimeMessage,
        })
    }

    ctx.JSON(200, gin.H{
        "messages": responseData,
    })
}
