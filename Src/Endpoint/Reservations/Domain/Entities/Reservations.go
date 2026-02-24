package entities

import "time"

type Reservation struct {
	ID              int       `json:"idReservation"`
	StudentID       int       `json:"studentId"`
	TeacherID       int       `json:"teacherId"`
	ReservationDate time.Time `json:"reservationDate"`
	ReservationTime string    `json:"reservationTime"`
	Attendance      *bool     `json:"attendance"`
	Topic           string    `json:"topic"`
}
