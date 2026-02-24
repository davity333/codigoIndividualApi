package application

import (
	entities "chat/Src/Endpoint/Class/Domain/Entities"
	repository "chat/Src/Endpoint/Class/Domain/Repository"
)

type UpdateClassUseCase struct {
	repository repository.IClass
}

func NewUpdateClassUseCase(repo repository.IClass) *UpdateClassUseCase {
	return &UpdateClassUseCase{
		repository: repo,
	}
}

func (u *UpdateClassUseCase) Execute(class *entities.Class) error {
	return u.repository.UpdateClass(class)
}
