package controller

import (
	application "chat/Src/Endpoint/Enrollment/Application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEnrollmentsByStudentIDController struct {
	usecase *application.GetEnrollmentsByStudentIDUseCase
}

func NewGetEnrollmentsByStudentIDController(usecase *application.GetEnrollmentsByStudentIDUseCase) *GetEnrollmentsByStudentIDController {
	return &GetEnrollmentsByStudentIDController{
		usecase: usecase,
	}
}

func (ctrl *GetEnrollmentsByStudentIDController) GetEnrollmentsByStudentID(c *gin.Context) {
	studentIDStr := c.Param("studentId")
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid studentId"})
		return
	}

	enrollments, err := ctrl.usecase.Execute(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}
