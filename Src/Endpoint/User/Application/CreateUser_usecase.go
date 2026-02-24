package application

import (
	entities "chat/Src/Endpoint/User/Domain/Entities"
	repository "chat/Src/Endpoint/User/Domain/Repository"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	usecase repository.IUser
}

func NewCreateUserUseCase(usecase repository.IUser) *CreateUserUseCase {
	return &CreateUserUseCase{
		usecase: usecase,
	}
}

func (c *CreateUserUseCase) ExecuteCreate(id int, username string, password string, email string, firstName string, lastName string, role string) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := entities.NewUser(id, username, email, string(hashedPassword), firstName, lastName, role)

	err = c.usecase.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
