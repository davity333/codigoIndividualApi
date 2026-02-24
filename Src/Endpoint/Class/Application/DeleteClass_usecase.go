package application

import repository "chat/Src/Endpoint/Class/Domain/Repository"

type DeleteClassUseCase struct {
	repository repository.IClass
}

func NewDeleteClassUseCase(repo repository.IClass) *DeleteClassUseCase {
	return &DeleteClassUseCase{
		repository: repo,
	}
}

func (u *DeleteClassUseCase) Execute(classID int64) error {
	return u.repository.DeleteClass(classID)
}
