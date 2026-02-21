package controller

import (
	application "chat/Src/Endpoint/Message/Application"
	"github.com/gin-gonic/gin"
	"strconv"
)


type DeleteMessageController struct {
	usecase application.DeleteMessageUseCase
}

func NewDeleteMessageController(usecase application.DeleteMessageUseCase) *DeleteMessageController {
	return &DeleteMessageController{
		usecase: usecase,
	}
}

func (c *DeleteMessageController) DeleteMessageController(ctx *gin.Context) {
    messageId := ctx.Param("messageId")
    messageIdInt, err := strconv.Atoi(messageId)

    if err != nil {
        ctx.JSON(400, gin.H{
            "error": "Error 400 - Solicitud incorrecta, el messageId debe ser un número entero",
            "Detail": err.Error(),
        })
        return
    }

    err = c.usecase.Execute(messageIdInt)
    if err != nil {
        ctx.JSON(500, gin.H{
            "error": "Error 500 - Error interno del servidor",
            "Detail": err.Error(),
        })
        return
    }

    ctx.JSON(200, gin.H{
        "message": "Mensaje eliminado exitosamente",
    })
}
