package repository

import entities "chat/Src/Endpoint/User/Domain/Entities"

type IUser interface {
	GetUserByName(username string) ([]*entities.User, error)
	GetAllUsers() ([]*entities.User, error)
	LoginUser(email, password string) (*entities.User, error)
	CreateUser(user *entities.User) error
	GetTeacherByID(userID string) (*entities.User, error)
}
