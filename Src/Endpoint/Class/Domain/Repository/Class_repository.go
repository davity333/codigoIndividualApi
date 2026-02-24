package repository

import (
	entities "chat/Src/Endpoint/Class/Domain/Entities"
)

type IClass interface {
	GetAllClasses() ([]entities.Class, error)
	GetClassByID(classID int64) (*entities.Class, error)
	GetClassesByTeacherID(teacherID int) ([]entities.Class, error)
	CreateClass(class *entities.Class) (*entities.Class, error)
	UpdateClass(class *entities.Class) error
	DeleteClass(classID int64) error
}
