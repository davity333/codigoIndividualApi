package controller

import (
    application "chat/Src/Endpoint/User/Application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GetTeacherByIDController struct {
    app *application.GetTeacherByIDUseCase
}

func NewGetTeacherByIDController(app *application.GetTeacherByIDUseCase) *GetTeacherByIDController {
    return &GetTeacherByIDController{
        app: app,
    }
}

func (c *GetTeacherByIDController) Handle(ctx *gin.Context) {
    id := ctx.Param("id")

    user, err := c.app.Execute(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": user,
    })
}
