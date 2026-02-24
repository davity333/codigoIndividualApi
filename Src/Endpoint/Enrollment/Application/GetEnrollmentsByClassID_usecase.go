package application

import (
	entities "chat/Src/Endpoint/Enrollment/Domain/Entities"
	repository "chat/Src/Endpoint/Enrollment/Domain/Repository"
)

type GetEnrollmentsByClassIDUseCase struct {
	repository repository.IEnrollment
}

func NewGetEnrollmentsByClassIDUseCase(repo repository.IEnrollment) *GetEnrollmentsByClassIDUseCase {
	return &GetEnrollmentsByClassIDUseCase{
		repository: repo,
	}
}

func (u *GetEnrollmentsByClassIDUseCase) Execute(classID int64) ([]entities.Enrollment, error) {
	return u.repository.GetEnrollmentsByClassID(classID)
}
