package controller

import (
	application "chat/Src/Endpoint/Class/Application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllClassesController struct {
	usecase *application.GetAllClassesUseCase
}

func NewGetAllClassesController(usecase *application.GetAllClassesUseCase) *GetAllClassesController {
	return &GetAllClassesController{
		usecase: usecase,
	}
}

func (ctrl *GetAllClassesController) GetAllClasses(c *gin.Context) {
	classes, err := ctrl.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}
