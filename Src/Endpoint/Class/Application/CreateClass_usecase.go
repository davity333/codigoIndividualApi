package application

import (
	entities "chat/Src/Endpoint/Class/Domain/Entities"
	repository "chat/Src/Endpoint/Class/Domain/Repository"
)

type CreateClassUseCase struct {
	repository repository.IClass
}

func NewCreateClassUseCase(repo repository.IClass) *CreateClassUseCase {
	return &CreateClassUseCase{
		repository: repo,
	}
}

func (u *CreateClassUseCase) Execute(class *entities.Class) (*entities.Class, error) {
	return u.repository.CreateClass(class)
}
