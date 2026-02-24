package controller

import (
	application "chat/Src/Endpoint/Class/Application"
	entities "chat/Src/Endpoint/Class/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateClassController struct {
	usecase *application.UpdateClassUseCase
}

func NewUpdateClassController(usecase *application.UpdateClassUseCase) *UpdateClassController {
	return &UpdateClassController{
		usecase: usecase,
	}
}

func (ctrl *UpdateClassController) UpdateClass(c *gin.Context) {
	var class entities.Class
	if err := c.BindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := ctrl.usecase.Execute(&class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}
