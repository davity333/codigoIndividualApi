package application

import (
	entities "chat/Src/Endpoint/Class/Domain/Entities"
	repository "chat/Src/Endpoint/Class/Domain/Repository"
)

type GetClassesByTeacherIDUseCase struct {
	repository repository.IClass
}

func NewGetClassesByTeacherIDUseCase(repo repository.IClass) *GetClassesByTeacherIDUseCase {
	return &GetClassesByTeacherIDUseCase{
		repository: repo,
	}
}

func (u *GetClassesByTeacherIDUseCase) Execute(teacherID int) ([]entities.Class, error) {
	return u.repository.GetClassesByTeacherID(teacherID)
}
