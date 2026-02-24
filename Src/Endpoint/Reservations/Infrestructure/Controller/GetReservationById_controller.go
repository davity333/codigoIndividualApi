package controller

import (
	application "chat/Src/Endpoint/Reservations/Application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetReservationByIDController struct {
	usecase *application.GetReservationByIDUseCase
}

func NewGetReservationByIDController(usecase *application.GetReservationByIDUseCase) *GetReservationByIDController {
	return &GetReservationByIDController{
		usecase: usecase,
	}
}

func (c *GetReservationByIDController) GetReservationByID(ctx *gin.Context) {
	id := ctx.Param("reservationId")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Error 400 - reservationId debe ser un número entero", "Detail": err.Error()})
		return
	}

	reservation, err := c.usecase.Execute(idInt)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Error 404 - Reservación no encontrada", "Detail": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": formatReservation(reservation)})
}
