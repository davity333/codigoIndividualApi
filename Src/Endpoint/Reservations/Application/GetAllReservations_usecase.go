package application

import (
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"
	repository "chat/Src/Endpoint/Reservations/Domain/Repository"
)

type GetAllReservationsUseCase struct {
	repository repository.IReservation
}

func NewGetAllReservationsUseCase(repository repository.IReservation) *GetAllReservationsUseCase {
	return &GetAllReservationsUseCase{
		repository: repository,
	}
}

func (uc *GetAllReservationsUseCase) Execute() ([]*entities.Reservation, error) {
	reservations, err := uc.repository.GetAllReservations()
	if err != nil {
		return nil, err
	}
	return reservations, nil
}
