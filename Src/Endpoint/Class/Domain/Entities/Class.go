package entities

import "time"

type Class struct {
	ID          int64     `json:"id"`
	TeacherID   int       `json:"teacherId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ClassDate   time.Time `json:"classDate"`
	StartTime   string    `json:"startTime"`
	EndTime     string    `json:"endTime"`
	Capacity    int       `json:"capacity"`
	Status      string    `json:"status"`
}
