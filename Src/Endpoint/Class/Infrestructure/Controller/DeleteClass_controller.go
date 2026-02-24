package controller

import (
	application "chat/Src/Endpoint/Class/Application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteClassController struct {
	usecase *application.DeleteClassUseCase
}

func NewDeleteClassController(usecase *application.DeleteClassUseCase) *DeleteClassController {
	return &DeleteClassController{
		usecase: usecase,
	}
}

func (ctrl *DeleteClassController) DeleteClass(c *gin.Context) {
	classIDStr := c.Param("classId")
	classID, err := strconv.ParseInt(classIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid classId"})
		return
	}

	err = ctrl.usecase.Execute(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
