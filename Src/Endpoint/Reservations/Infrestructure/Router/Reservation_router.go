package router

import (
	controller "chat/Src/Endpoint/Reservations/Infrestructure/Controller"

	"github.com/gin-gonic/gin"
)

func ReservationRouter(
	g *gin.Engine,
	getAllReservations *controller.GetAllReservationsController,
	getReservationByID *controller.GetReservationByIDController,
	createReservation *controller.CreateReservationController,
	updateReservation *controller.UpdateReservationController,
	deleteReservation *controller.DeleteReservationController,
	getReservationsByStudentID *controller.GetReservationsByStudentIDController,
) {
	reservationGroup := g.Group("/api/v1/reservations")
	{
		reservationGroup.GET("/getAll", getAllReservations.GetAllReservations)
		reservationGroup.GET("/:reservationId", getReservationByID.GetReservationByID)
		reservationGroup.GET("/student/:studentId", getReservationsByStudentID.GetReservationsByStudentID)
		reservationGroup.POST("/create", createReservation.CreateReservation)
		reservationGroup.PUT("/update/:reservationId", updateReservation.UpdateReservation)
		reservationGroup.DELETE("/delete/:reservationId", deleteReservation.DeleteReservation)
	}
}
