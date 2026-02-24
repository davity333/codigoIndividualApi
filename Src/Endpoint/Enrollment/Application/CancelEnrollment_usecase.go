package application

import repository "chat/Src/Endpoint/Enrollment/Domain/Repository"

type CancelEnrollmentUseCase struct {
	repository repository.IEnrollment
}

func NewCancelEnrollmentUseCase(repo repository.IEnrollment) *CancelEnrollmentUseCase {
	return &CancelEnrollmentUseCase{
		repository: repo,
	}
}

func (u *CancelEnrollmentUseCase) Execute(enrollmentID int64) error {
	return u.repository.CancelEnrollment(enrollmentID)
}
