package controller

import (
	application "chat/Src/Endpoint/Reservations/Application"
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateReservationController struct {
	usecase *application.UpdateReservationUseCase
}

func NewUpdateReservationController(usecase *application.UpdateReservationUseCase) *UpdateReservationController {
	return &UpdateReservationController{
		usecase: usecase,
	}
}

func (c *UpdateReservationController) UpdateReservation(ctx *gin.Context) {
	id := ctx.Param("reservationId")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Error 400 - reservationId debe ser un número entero", "Detail": err.Error()})
		return
	}

	var reservation *entities.Reservation
	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(400, gin.H{"error": "Error 400 - Solicitud incorrecta", "Detail": err.Error()})
		return
	}

	err = c.usecase.Execute(idInt, reservation)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Error 404 - Reservación no encontrada", "Detail": err.Error()})
		return
	}

	reservation.ID = idInt
	ctx.JSON(200, gin.H{"data": formatReservation(reservation)})
}
