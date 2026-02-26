package controller

import (
	application "chat/Src/Endpoint/User/Application"
	"net/http"
    "fmt"
	"github.com/gin-gonic/gin"
)

type GetUserByNameController struct {
	usecase *application.GetUserByNameUseCase
}

func NewGetUserByNameController(usecase *application.GetUserByNameUseCase) *GetUserByNameController {
	return &GetUserByNameController{
		usecase: usecase,
	}
}

func (g *GetUserByNameController) GetByUsername(ctx *gin.Context) {
    username := ctx.Param("username")

    if username == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error":  "Error 400 - Nombre vacío",
            "detail": "El nombre del usuario no puede estar vacío",
        })
        return
    }

    users, err := g.usecase.Execute(username)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error":  "Error 404 - Usuario no encontrado",
            "detail": fmt.Sprintf("El usuario %s no existe", username),
        })
        return
    }

    var responseData []gin.H
    for _, user := range users {
        userData := gin.H{
            "type": "users",
            "idUser": user.ID,
            "attributes": gin.H{
                "username":  user.Username,
                "email":     user.Email,
                "firstName": user.FirstName,
                "lastName":  user.LastName,
                "role":      user.Role,
            },
        }
        responseData = append(responseData, userData)
    }

    ctx.JSON(http.StatusOK, responseData)
}
