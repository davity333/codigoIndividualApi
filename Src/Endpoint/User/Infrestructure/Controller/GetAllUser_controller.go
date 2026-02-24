package controller

import (
	application "chat/Src/Endpoint/User/Application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllUserController struct {
	usecase *application.GetAllUsersUseCase
}

func NewGetAllUserController(usecase *application.GetAllUsersUseCase) *GetAllUserController {
	return &GetAllUserController{
		usecase: usecase,
	}
}

func (g *GetAllUserController) GetUser(c *gin.Context) {
	user, err := g.usecase.Execute()

	if err != nil {
		fmt.Println("Error al obtener usuarios:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar los usuarios"})
		return
	}

	fmt.Println("Usuarios obtenidos:", user)

	var usersData []gin.H
	for _, users := range user {
		userData := gin.H{
			"type":   "Users",
			"idUser": users.ID,
			"attributes": gin.H{
				"username":  users.Username,
				"Email":     users.Email,
				"firstName": users.FirstName,
				"lastName":  users.LastName,
				"role":      users.Role,
			},
		}
		usersData = append(usersData, userData)
	}

	response := gin.H{
		"data": usersData,
	}
	c.JSON(http.StatusOK, response)
}
