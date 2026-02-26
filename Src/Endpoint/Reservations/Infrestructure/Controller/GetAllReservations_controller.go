package controller

import (
	application "chat/Src/Endpoint/Reservations/Application"

	"github.com/gin-gonic/gin"
)

type GetAllReservationsController struct {
	usecase *application.GetAllReservationsUseCase
}

func NewGetAllReservationsController(usecase *application.GetAllReservationsUseCase) *GetAllReservationsController {
	return &GetAllReservationsController{
		usecase: usecase,
	}
}

func (c *GetAllReservationsController) GetAllReservations(ctx *gin.Context) {
    reservations, err := c.usecase.Execute()
    if err != nil {
        ctx.JSON(500, gin.H{
            "error":  "Error 500 - Error interno del servidor",
            "detail": err.Error(),
        })
        return
    }

    // Si no hay reservaciones
    if len(reservations) == 0 {
        ctx.JSON(200, gin.H{
            "message": "No hay ninguna reservación",
            "reservations": []gin.H{},
        })
        return
    }

    // Si sí hay reservaciones
    responseData := make([]gin.H, 0, len(reservations))
    for _, reservation := range reservations {
        responseData = append(responseData, formatReservation(reservation))
    }

    ctx.JSON(200, gin.H{
        "reservations": responseData,
    })
}
