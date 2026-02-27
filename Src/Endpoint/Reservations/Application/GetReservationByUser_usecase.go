package application

import (
    entities "chat/Src/Endpoint/Reservations/Domain/Entities"
    repository "chat/Src/Endpoint/Reservations/Domain/Repository"
)

type GetReservationsByStudentIDUseCase struct {
    repo repository.IReservation
}

func NewGetReservationsByStudentIDUseCase(repo repository.IReservation) *GetReservationsByStudentIDUseCase {
    return &GetReservationsByStudentIDUseCase{repo: repo}
}

func (uc *GetReservationsByStudentIDUseCase) Execute(studentID int) ([]*entities.Reservation, error) {
    return uc.repo.GetReservationsByStudentID(studentID)
}
