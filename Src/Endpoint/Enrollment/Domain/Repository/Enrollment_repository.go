package repository

import (
	entities "chat/Src/Endpoint/Enrollment/Domain/Entities"
)

type IEnrollment interface {
	GetAllEnrollments() ([]entities.Enrollment, error)
	GetEnrollmentByID(enrollmentID int64) (*entities.Enrollment, error)
	GetEnrollmentsByClassID(classID int64) ([]entities.Enrollment, error)
	GetEnrollmentsByStudentID(studentID int) ([]entities.Enrollment, error)
	CreateEnrollment(enrollment *entities.Enrollment) (*entities.Enrollment, error)
	CancelEnrollment(enrollmentID int64) error
	CompleteEnrollment(enrollmentID int64) error
}
