package application

import (
	auth "chat/Src/Auth"
	entities "chat/Src/Endpoint/User/Domain/Entities"
	repository "chat/Src/Endpoint/User/Domain/Repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	usecase repository.IUser
}

func NewLoginUserUseCase(usecase repository.IUser) *LoginUserUseCase {
	return &LoginUserUseCase{
		usecase: usecase,
	}
}

func (c *LoginUserUseCase) ExecuteLogin(email string, password string) (*entities.User, string, error) {
	user, err := c.usecase.LoginUser(email, password)
	if err != nil || user == nil {
		return nil, "", fmt.Errorf("usuario no encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", fmt.Errorf("contraseña incorrecta tilin")
	}

	token, err := auth.GenerateToken(email)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
