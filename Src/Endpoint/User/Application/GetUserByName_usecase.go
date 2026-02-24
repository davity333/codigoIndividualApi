package application

import (
	entities "chat/Src/Endpoint/User/Domain/Entities"
	repository "chat/Src/Endpoint/User/Domain/Repository"
)

type GetUserByNameUseCase struct {
	usecase repository.IUser
}

func NewGetUserByNameUseCase(usecase repository.IUser) *GetUserByNameUseCase {
	return &GetUserByNameUseCase{
		usecase: usecase,
	}
}

func (uc *GetUserByNameUseCase) Execute(username string) ([]*entities.User, error) {
	user, err := uc.usecase.GetUserByName(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
