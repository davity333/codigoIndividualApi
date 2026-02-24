package application

import (
	entities "chat/Src/Endpoint/Class/Domain/Entities"
	repository "chat/Src/Endpoint/Class/Domain/Repository"
)

type GetClassByIDUseCase struct {
	repository repository.IClass
}

func NewGetClassByIDUseCase(repo repository.IClass) *GetClassByIDUseCase {
	return &GetClassByIDUseCase{
		repository: repo,
	}
}

func (u *GetClassByIDUseCase) Execute(classID int64) (*entities.Class, error) {
	return u.repository.GetClassByID(classID)
}
