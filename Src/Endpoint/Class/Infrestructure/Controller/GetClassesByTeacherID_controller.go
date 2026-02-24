package controller

import (
	application "chat/Src/Endpoint/Class/Application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetClassesByTeacherIDController struct {
	usecase *application.GetClassesByTeacherIDUseCase
}

func NewGetClassesByTeacherIDController(usecase *application.GetClassesByTeacherIDUseCase) *GetClassesByTeacherIDController {
	return &GetClassesByTeacherIDController{
		usecase: usecase,
	}
}

func (ctrl *GetClassesByTeacherIDController) GetClassesByTeacherID(c *gin.Context) {
	teacherIDStr := c.Param("teacherId")
	teacherID, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacherId"})
		return
	}

	classes, err := ctrl.usecase.Execute(teacherID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classes)
}
