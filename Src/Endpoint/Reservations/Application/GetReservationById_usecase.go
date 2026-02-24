package application

import (
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"
	repository "chat/Src/Endpoint/Reservations/Domain/Repository"
)

type GetReservationByIDUseCase struct {
	repository repository.IReservation
}

func NewGetReservationByIDUseCase(repository repository.IReservation) *GetReservationByIDUseCase {
	return &GetReservationByIDUseCase{
		repository: repository,
	}
}

func (uc *GetReservationByIDUseCase) Execute(id int) (*entities.Reservation, error) {
	reservation, err := uc.repository.GetReservationByID(id)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}
