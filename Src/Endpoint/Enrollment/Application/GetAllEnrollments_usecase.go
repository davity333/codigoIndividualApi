package application

import (
	entities "chat/Src/Endpoint/Enrollment/Domain/Entities"
	repository "chat/Src/Endpoint/Enrollment/Domain/Repository"
)

type GetAllEnrollmentsUseCase struct {
	repository repository.IEnrollment
}

func NewGetAllEnrollmentsUseCase(repo repository.IEnrollment) *GetAllEnrollmentsUseCase {
	return &GetAllEnrollmentsUseCase{
		repository: repo,
	}
}

func (u *GetAllEnrollmentsUseCase) Execute() ([]entities.Enrollment, error) {
	return u.repository.GetAllEnrollments()
}
