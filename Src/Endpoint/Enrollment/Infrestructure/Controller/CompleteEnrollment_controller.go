package controller

import (
	application "chat/Src/Endpoint/Enrollment/Application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompleteEnrollmentController struct {
	usecase *application.CompleteEnrollmentUseCase
}

func NewCompleteEnrollmentController(usecase *application.CompleteEnrollmentUseCase) *CompleteEnrollmentController {
	return &CompleteEnrollmentController{
		usecase: usecase,
	}
}

func (ctrl *CompleteEnrollmentController) CompleteEnrollment(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment completed successfully"})
}
