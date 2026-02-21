package controller

import (
	application "chat/Src/Endpoint/User/Application"
	entities "chat/Src/Endpoint/User/Domain/Entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserController struct {
	usecase *application.CreateUserUseCase
}

func NewCreateUserController(usecase *application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{
		usecase: usecase,
	}
}

func (c *CreateUserController) CreateUser(ctx *gin.Context) {
	var user *entities.User

	fmt.Println("Datos del cuerpo de la solicitud:", ctx.Request.Body)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Error 404 - Solicitud incorrecta, el cuerpo de la solicitud no es válido",
			"Detail": err.Error()})
		return
	}

	user, err := c.usecase.ExecuteCreate(user.ID, user.Username, user.Password, user.Email, user.FirstName, user.LastName, user.Role)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error 500 - Error interno del servidor, no se pudo crear el usuario",
			"Detail": err.Error()})
		return
	}

	response := gin.H{
		"data": gin.H{
			"type": "services",
			"idUser": user.ID,
			"attributes": gin.H{
				"username": user.Username,
				"email":     user.Email,
				"firstName": user.FirstName,
				"lastName": user.LastName,
				"role": user.Role,
			},
		}}
		ctx.JSON(http.StatusCreated, response)
}
