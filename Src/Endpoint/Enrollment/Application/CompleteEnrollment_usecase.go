package application

import repository "chat/Src/Endpoint/Enrollment/Domain/Repository"

type CompleteEnrollmentUseCase struct {
	repository repository.IEnrollment
}

func NewCompleteEnrollmentUseCase(repo repository.IEnrollment) *CompleteEnrollmentUseCase {
	return &CompleteEnrollmentUseCase{
		repository: repo,
	}
}

func (u *CompleteEnrollmentUseCase) Execute(enrollmentID int64) error {
	return u.repository.CompleteEnrollment(enrollmentID)
}
