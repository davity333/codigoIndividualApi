package controller

import (
    application "chat/Src/Endpoint/Contacts/Application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GetContactByNameController struct {
    usecase *application.GetContactByNameUseCase
}

func NewGetContactByNameController(usecase *application.GetContactByNameUseCase) *GetContactByNameController {
    return &GetContactByNameController{usecase: usecase}
}

func (c *GetContactByNameController) Handle(ctx *gin.Context) {
    username := ctx.Param("username")

    user, err := c.usecase.Execute(username)
    if err != nil {
        switch err {
        case application.ErrEmptyUsername:
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        case application.ErrUserNotFound:
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        default:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "user found",
        "data":    user,
    })
}
