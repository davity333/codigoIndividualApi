package application

import (
	entities "chat/Src/Endpoint/User/Domain/Entities"
	repository "chat/Src/Endpoint/User/Domain/Repository"
)

type GetTeacherByIDUseCase struct {
	usecase repository.IUser
}

func NewGetTeacherByIDUseCase(usecase repository.IUser) *GetTeacherByIDUseCase {
	return &GetTeacherByIDUseCase{
		usecase: usecase,
	}
}

func (uc *GetTeacherByIDUseCase) Execute(userID string) (*entities.User, error) {
	user, err := uc.usecase.GetTeacherByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}