package controller

import (
    application "chat/Src/Endpoint/Class/Application"
    "github.com/gin-gonic/gin"
)

type GetClassesByDateController struct {
    usecase *application.GetClassesByDateUseCase
}

func NewGetClassesByDateController(usecase *application.GetClassesByDateUseCase) *GetClassesByDateController {
    return &GetClassesByDateController{usecase: usecase}
}

func (c *GetClassesByDateController) GetClassesByDate(ctx *gin.Context) {
    date := ctx.Param("date") // formato: 2026-02-26

    classes, err := c.usecase.Execute(date)
    if err != nil {
        ctx.JSON(500, gin.H{
            "error":  "Error interno del servidor",
            "detail": err.Error(),
        })
        return
    }

    ctx.JSON(200, gin.H{
        "classes": classes,
    })
}
