package controller

import (
	application "chat/Src/Endpoint/Reservations/Application"
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"

	"github.com/gin-gonic/gin"
)

type CreateReservationController struct {
	usecase *application.CreateReservationUseCase
}

func NewCreateReservationController(usecase *application.CreateReservationUseCase) *CreateReservationController {
	return &CreateReservationController{
		usecase: usecase,
	}
}

func (c *CreateReservationController) CreateReservation(ctx *gin.Context) {
	var reservation *entities.Reservation

	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(400, gin.H{"error": "Error 400 - Solicitud incorrecta", "Detail": err.Error()})
		return
	}

	err := c.usecase.Execute(reservation)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error 500 - Error interno del servidor", "Detail": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"data": formatReservation(reservation)})
}
