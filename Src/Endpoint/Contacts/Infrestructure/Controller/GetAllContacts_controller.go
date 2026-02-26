package controller

import (
    application "chat/Src/Endpoint/Contacts/Application"
    "chat/Src/Endpoint/Contacts/Domain/Entities"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

type GetAllContactsController struct {
    usecase *application.GetAllContactsUseCase
}

func NewGetAllContactsController(usecase *application.GetAllContactsUseCase) *GetAllContactsController {
    return &GetAllContactsController{
        usecase: usecase,
    }
}

func (c *GetAllContactsController) Handle(ctx *gin.Context) {
    userIDParam := ctx.Param("userId")

    var userID int
    _, err := fmt.Sscanf(userIDParam, "%d", &userID)
    if err != nil || userID <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId inválido"})
        return
    }

    contacts, err := c.usecase.Execute(userID)
if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
}

if len(contacts) == 0 {
    ctx.JSON(http.StatusOK, gin.H{
        "message":  "No tienes ningún contacto agregado",
        "contacts": []entities.ContactResponse{},
    })
    return
}

ctx.JSON(http.StatusOK, gin.H{
    "message":  "Contactos obtenidos correctamente",
    "contacts": contacts,
})

}
