package controller

import (
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"

	"github.com/gin-gonic/gin"
)

func formatReservation(reservation *entities.Reservation) gin.H {
	return gin.H{
		"idReservation":   reservation.ID,
		"studentId":       reservation.StudentID,
		"classId":         reservation.ClassID,
		"reservationDate": reservation.ReservationDate.Format("2006-01-02"),
		"attendance":      reservation.Attendance,
	}
}
