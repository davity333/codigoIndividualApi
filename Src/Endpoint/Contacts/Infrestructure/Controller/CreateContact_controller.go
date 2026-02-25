package controller

import (
    application "chat/Src/Endpoint/Contacts/Application"
    entities "chat/Src/Endpoint/Contacts/Domain/Entities"
    "github.com/gin-gonic/gin"
    "net/http"
)

type CreateContactController struct {
    usecase *application.CreateContactUseCase
}

func NewCreateContactController(usecase *application.CreateContactUseCase) *CreateContactController {
    return &CreateContactController{
        usecase: usecase,
    }
}

func (c *CreateContactController) Handle(ctx *gin.Context) {
    var req entities.Contact

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
        return
    }

    if req.UserID <= 0 || req.ContactID <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "IDs inválidos"})
        return
    }

    err := c.usecase.Execute(req)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Contacto agregado correctamente",
    })
}
