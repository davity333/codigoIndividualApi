package controller

import (
	application "chat/Src/Endpoint/User/Application"
	entities "chat/Src/Endpoint/User/Domain/Entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUserUseCase struct {
	usecase *application.LoginUserUseCase
}

func NewLoginUserUseCase(usecase *application.LoginUserUseCase) *LoginUserUseCase {
	return &LoginUserUseCase{
		usecase: usecase,
	}
}

func (c *LoginUserUseCase) LoginUser(ctx *gin.Context) {
	var user *entities.User
	fmt.Println("Datos del cuerpo de la solicitud:", ctx.Request.Body)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error":  "Error 404 - Solicitud incorrecta, el cuerpo de la solicitud no es válido",
			"Detail": err.Error()})
		return
	}

	user, token, err := c.usecase.ExecuteLogin(user.Email, user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":  "Error 500 - Error interno del servidor, no se pudo iniciar sesión",
			"Detail": err.Error()})
		return
	}

	response := gin.H{
		"data": gin.H{
			"type":   "user",
			"idUser": user.ID,
			"attributes": gin.H{
				"username":  user.Username,
				"email":     user.Email,
				"firstName": user.FirstName,
				"lastName":  user.LastName,
				"role":      user.Role,
			},
			"token": token,
		}}
	ctx.JSON(http.StatusOK, response)
}
