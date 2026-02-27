package repository

import entities "chat/Src/Endpoint/Reservations/Domain/Entities"

type IReservation interface {
	GetAllReservations() ([]*entities.Reservation, error)
	GetReservationByID(id int) (*entities.Reservation, error)
	CreateReservation(reservation *entities.Reservation) error
	UpdateReservation(id int, reservation *entities.Reservation) error
	DeleteReservation(id int) error
	GetReservationsByStudentID(studentID int) ([]*entities.Reservation, error)
}
