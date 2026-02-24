package application

import repository "chat/Src/Endpoint/Reservations/Domain/Repository"

type DeleteReservationUseCase struct {
	repository repository.IReservation
}

func NewDeleteReservationUseCase(repository repository.IReservation) *DeleteReservationUseCase {
	return &DeleteReservationUseCase{
		repository: repository,
	}
}

func (uc *DeleteReservationUseCase) Execute(id int) error {
	return uc.repository.DeleteReservation(id)
}
