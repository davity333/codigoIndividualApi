package controller

import (
    "strconv"

    application "chat/Src/Endpoint/Reservations/Application"
    "github.com/gin-gonic/gin"
)

type GetReservationsByStudentIDController struct {
    usecase *application.GetReservationsByStudentIDUseCase
}

func NewGetReservationsByStudentIDController(usecase *application.GetReservationsByStudentIDUseCase) *GetReservationsByStudentIDController {
    return &GetReservationsByStudentIDController{usecase: usecase}
}

func (c *GetReservationsByStudentIDController) GetReservationsByStudentID(ctx *gin.Context) {

    studentIDStr := ctx.Param("studentId")
    studentID, err := strconv.Atoi(studentIDStr)
    if err != nil {
        ctx.JSON(400, gin.H{"error": "studentId inválido"})
        return
    }

    reservations, err := c.usecase.Execute(studentID)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Error interno del servidor", "detail": err.Error()})
        return
    }

    ctx.JSON(200, gin.H{"reservations": reservations})
}
