package controller

import (
    application "chat/Src/Endpoint/Contacts/Application"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

type DeleteContactController struct {
    usecase *application.DeleteContactUseCase
}

func NewDeleteContactController(usecase *application.DeleteContactUseCase) *DeleteContactController {
    return &DeleteContactController{
        usecase: usecase,
    }
}

func (c *DeleteContactController) Handle(ctx *gin.Context) {
    userIDParam := ctx.Param("userId")
    contactIDParam := ctx.Param("contactId")

    var userID, contactID int

    _, err := fmt.Sscanf(userIDParam, "%d", &userID)
    if err != nil || userID <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId inválido"})
        return
    }

    _, err = fmt.Sscanf(contactIDParam, "%d", &contactID)
    if err != nil || contactID <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "contactId inválido"})
        return
    }

    err = c.usecase.Execute(userID, contactID)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Contacto eliminado correctamente",
    })
}
