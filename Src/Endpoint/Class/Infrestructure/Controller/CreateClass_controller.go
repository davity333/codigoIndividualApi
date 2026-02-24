package controller

import (
	application "chat/Src/Endpoint/Class/Application"
	entities "chat/Src/Endpoint/Class/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateClassController struct {
	usecase *application.CreateClassUseCase
}

func NewCreateClassController(usecase *application.CreateClassUseCase) *CreateClassController {
	return &CreateClassController{
		usecase: usecase,
	}
}

func (ctrl *CreateClassController) CreateClass(c *gin.Context) {
	var class entities.Class
	if err := c.BindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdClass, err := ctrl.usecase.Execute(&class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdClass)
}
