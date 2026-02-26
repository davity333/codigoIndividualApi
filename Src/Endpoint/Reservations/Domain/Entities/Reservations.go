package entities

import "time"

type Reservation struct {
	ID              int       `json:"idReservation"`
	StudentID       int       `json:"studentId"`
	ClassID       int       `json:"classId"`
	ReservationDate time.Time `json:"reservationDate"`
	Attendance      *bool     `json:"attendance"`
}
