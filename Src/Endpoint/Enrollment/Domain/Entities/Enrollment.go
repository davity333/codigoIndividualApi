package entities

import "time"

type Enrollment struct {
	ID         int64     `json:"id"`
	ClassID    int64     `json:"classId"`
	StudentID  int       `json:"studentId"`
	EnrolledAt time.Time `json:"enrolledAt"`
	Status     string    `json:"status"` // "Activa", "Cancelada", "Completada"
}
