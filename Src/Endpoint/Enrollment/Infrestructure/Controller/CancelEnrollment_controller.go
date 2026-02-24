package controller

import (
	application "chat/Src/Endpoint/Enrollment/Application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CancelEnrollmentController struct {
	usecase *application.CancelEnrollmentUseCase
}

func NewCancelEnrollmentController(usecase *application.CancelEnrollmentUseCase) *CancelEnrollmentController {
	return &CancelEnrollmentController{
		usecase: usecase,
	}
}

func (ctrl *CancelEnrollmentController) CancelEnrollment(c *gin.Context) {
	enrollmentIDStr := c.Param("enrollmentId")
	enrollmentID, err := strconv.ParseInt(enrollmentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollmentId"})
		return
	}

	err = ctrl.usecase.Execute(enrollmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment cancelled successfully"})
}
