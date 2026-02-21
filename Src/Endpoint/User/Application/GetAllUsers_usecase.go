package application

import (
	entities "chat/Src/Endpoint/User/Domain/Entities"
	repository "chat/Src/Endpoint/User/Domain/Repository"
)

type GetAllUsersUseCase struct {
	usecase repository.IUser
}

func NewGetAllUsersUseCase(usecase repository.IUser) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{
		usecase: usecase,
	}
}

func (uc *GetAllUsersUseCase) Execute() ([]*entities.User, error) {
	users, err := uc.usecase.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}