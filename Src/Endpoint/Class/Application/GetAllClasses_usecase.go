package application

import (
	entities "chat/Src/Endpoint/Class/Domain/Entities"
	repository "chat/Src/Endpoint/Class/Domain/Repository"
)

type GetAllClassesUseCase struct {
	repository repository.IClass
}

func NewGetAllClassesUseCase(repo repository.IClass) *GetAllClassesUseCase {
	return &GetAllClassesUseCase{
		repository: repo,
	}
}

func (u *GetAllClassesUseCase) Execute() ([]entities.ClassWithTeacher, error) {
    return u.repository.GetAllClasses()
}

