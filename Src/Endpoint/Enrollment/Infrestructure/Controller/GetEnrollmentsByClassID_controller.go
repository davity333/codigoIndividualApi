package controller

import (
	application "chat/Src/Endpoint/Enrollment/Application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEnrollmentsByClassIDController struct {
	usecase *application.GetEnrollmentsByClassIDUseCase
}

func NewGetEnrollmentsByClassIDController(usecase *application.GetEnrollmentsByClassIDUseCase) *GetEnrollmentsByClassIDController {
	return &GetEnrollmentsByClassIDController{
		usecase: usecase,
	}
}

func (ctrl *GetEnrollmentsByClassIDController) GetEnrollmentsByClassID(c *gin.Context) {
	classIDStr := c.Param("classId")
	classID, err := strconv.ParseInt(classIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid classId"})
		return
	}

	enrollments, err := ctrl.usecase.Execute(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}
