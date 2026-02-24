package sql

import (
	config "chat/Src/Core"
	entities "chat/Src/Endpoint/Enrollment/Domain/Entities"
	"database/sql"
	"time"
)

type EnrollmentSQL struct {
	config *config.ConnMySQL
}

func NewEnrollmentSQL() (*EnrollmentSQL, error) {
	conn := config.GetDBPool()
	return &EnrollmentSQL{config: conn}, nil
}

func (e *EnrollmentSQL) GetAllEnrollments() ([]entities.Enrollment, error) {
	query := `SELECT idEnrollment, classId, studentId, enrolledAt, status FROM enrollments WHERE status = 'Activa'`
	rows, err := e.config.FetchRows(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []entities.Enrollment
	for rows.Next() {
		enrollment, err := scanEnrollment(rows)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}

func (e *EnrollmentSQL) GetEnrollmentByID(enrollmentID int64) (*entities.Enrollment, error) {
	query := `SELECT idEnrollment, classId, studentId, enrolledAt, status FROM enrollments WHERE idEnrollment = ?`
	row := e.config.QueryRow(query, enrollmentID)
	enrollment, err := scanEnrollmentRow(row)
	if err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (e *EnrollmentSQL) GetEnrollmentsByClassID(classID int64) ([]entities.Enrollment, error) {
	query := `SELECT idEnrollment, classId, studentId, enrolledAt, status FROM enrollments WHERE classId = ? AND status = 'Activa'`
	rows, err := e.config.FetchRows(query, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []entities.Enrollment
	for rows.Next() {
		enrollment, err := scanEnrollment(rows)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}

func (e *EnrollmentSQL) GetEnrollmentsByStudentID(studentID int) ([]entities.Enrollment, error) {
	query := `SELECT idEnrollment, classId, studentId, enrolledAt, status FROM enrollments WHERE studentId = ? AND status = 'Activa'`
	rows, err := e.config.FetchRows(query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []entities.Enrollment
	for rows.Next() {
		enrollment, err := scanEnrollment(rows)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}

func (e *EnrollmentSQL) CreateEnrollment(enrollment *entities.Enrollment) (*entities.Enrollment, error) {
	query := `INSERT INTO enrollments (classId, studentId, enrolledAt, status) VALUES (?, ?, ?, ?)`
	result, err := e.config.ExecutePreparedQuery(query, enrollment.ClassID, enrollment.StudentID, time.Now(), "Activa")
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	enrollment.ID = id
	enrollment.EnrolledAt = time.Now()
	enrollment.Status = "Activa"
	return enrollment, nil
}

func (e *EnrollmentSQL) CancelEnrollment(enrollmentID int64) error {
	query := `UPDATE enrollments SET status = 'Cancelada' WHERE idEnrollment = ?`
	_, err := e.config.ExecutePreparedQuery(query, enrollmentID)
	return err
}

func (e *EnrollmentSQL) CompleteEnrollment(enrollmentID int64) error {
	query := `UPDATE enrollments SET status = 'Completada' WHERE idEnrollment = ?`
	_, err := e.config.ExecutePreparedQuery(query, enrollmentID)
	return err
}

func scanEnrollment(rows *sql.Rows) (entities.Enrollment, error) {
	var enrollment entities.Enrollment
	err := rows.Scan(&enrollment.ID, &enrollment.ClassID, &enrollment.StudentID, &enrollment.EnrolledAt, &enrollment.Status)
	return enrollment, err
}

func scanEnrollmentRow(row *sql.Row) (entities.Enrollment, error) {
	var enrollment entities.Enrollment
	err := row.Scan(&enrollment.ID, &enrollment.ClassID, &enrollment.StudentID, &enrollment.EnrolledAt, &enrollment.Status)
	return enrollment, err
}
