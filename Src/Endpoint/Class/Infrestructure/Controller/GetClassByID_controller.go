package controller

import (
	application "chat/Src/Endpoint/Class/Application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetClassByIDController struct {
	usecase *application.GetClassByIDUseCase
}

func NewGetClassByIDController(usecase *application.GetClassByIDUseCase) *GetClassByIDController {
	return &GetClassByIDController{
		usecase: usecase,
	}
}

func (ctrl *GetClassByIDController) GetClassByID(c *gin.Context) {
	classIDStr := c.Param("classId")
	classID, err := strconv.ParseInt(classIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid classId"})
		return
	}

	class, err := ctrl.usecase.Execute(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, class)
}
