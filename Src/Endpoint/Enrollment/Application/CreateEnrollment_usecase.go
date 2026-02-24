package application

import (
	entities "chat/Src/Endpoint/Enrollment/Domain/Entities"
	repository "chat/Src/Endpoint/Enrollment/Domain/Repository"
)

type CreateEnrollmentUseCase struct {
	repository repository.IEnrollment
}

func NewCreateEnrollmentUseCase(repo repository.IEnrollment) *CreateEnrollmentUseCase {
	return &CreateEnrollmentUseCase{
		repository: repo,
	}
}

func (u *CreateEnrollmentUseCase) Execute(enrollment *entities.Enrollment) (*entities.Enrollment, error) {
	return u.repository.CreateEnrollment(enrollment)
}
