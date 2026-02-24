package controller

import (
	application "chat/Src/Endpoint/Enrollment/Application"
	entities "chat/Src/Endpoint/Enrollment/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEnrollmentController struct {
	usecase *application.CreateEnrollmentUseCase
}

func NewCreateEnrollmentController(usecase *application.CreateEnrollmentUseCase) *CreateEnrollmentController {
	return &CreateEnrollmentController{
		usecase: usecase,
	}
}

func (ctrl *CreateEnrollmentController) CreateEnrollment(c *gin.Context) {
	var enrollment entities.Enrollment
	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdEnrollment, err := ctrl.usecase.Execute(&enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdEnrollment)
}
