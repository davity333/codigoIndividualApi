package application

import (
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"
	repository "chat/Src/Endpoint/Reservations/Domain/Repository"
)

type UpdateReservationUseCase struct {
	repository repository.IReservation
}

func NewUpdateReservationUseCase(repository repository.IReservation) *UpdateReservationUseCase {
	return &UpdateReservationUseCase{
		repository: repository,
	}
}

func (uc *UpdateReservationUseCase) Execute(id int, reservation *entities.Reservation) error {
	return uc.repository.UpdateReservation(id, reservation)
}
