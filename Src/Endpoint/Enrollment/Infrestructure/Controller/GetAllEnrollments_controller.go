package controller

import (
	application "chat/Src/Endpoint/Enrollment/Application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllEnrollmentsController struct {
	usecase *application.GetAllEnrollmentsUseCase
}

func NewGetAllEnrollmentsController(usecase *application.GetAllEnrollmentsUseCase) *GetAllEnrollmentsController {
	return &GetAllEnrollmentsController{
		usecase: usecase,
	}
}

func (ctrl *GetAllEnrollmentsController) GetAllEnrollments(c *gin.Context) {
	enrollments, err := ctrl.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollments)
}
