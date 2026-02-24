package application

import (
	entities "chat/Src/Endpoint/Enrollment/Domain/Entities"
	repository "chat/Src/Endpoint/Enrollment/Domain/Repository"
)

type GetEnrollmentsByStudentIDUseCase struct {
	repository repository.IEnrollment
}

func NewGetEnrollmentsByStudentIDUseCase(repo repository.IEnrollment) *GetEnrollmentsByStudentIDUseCase {
	return &GetEnrollmentsByStudentIDUseCase{
		repository: repo,
	}
}

func (u *GetEnrollmentsByStudentIDUseCase) Execute(studentID int) ([]entities.Enrollment, error) {
	return u.repository.GetEnrollmentsByStudentID(studentID)
}
