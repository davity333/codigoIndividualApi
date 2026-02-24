package application

import (
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"
	repository "chat/Src/Endpoint/Reservations/Domain/Repository"
)

type CreateReservationUseCase struct {
	repository repository.IReservation
}

func NewCreateReservationUseCase(repository repository.IReservation) *CreateReservationUseCase {
	return &CreateReservationUseCase{
		repository: repository,
	}
}

func (uc *CreateReservationUseCase) Execute(reservation *entities.Reservation) error {
	return uc.repository.CreateReservation(reservation)
}
