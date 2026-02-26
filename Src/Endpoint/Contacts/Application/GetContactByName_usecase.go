package application

import (
    "errors"
    "strings"

    userEntities "chat/Src/Endpoint/User/Domain/Entities"
    repository "chat/Src/Endpoint/Contacts/Domain/Repository"
)

var (
    ErrEmptyUsername   = errors.New("username cannot be empty")
    ErrUserNotFound    = errors.New("user not found")
)

type GetContactByNameUseCase struct {
    repo repository.IContacts
}

func NewGetContactByNameUseCase(repo repository.IContacts) *GetContactByNameUseCase {
    return &GetContactByNameUseCase{repo: repo}
}

func (uc *GetContactByNameUseCase) Execute(username string) (*userEntities.User, error) {
    // 1. Validación de entrada
    username = strings.TrimSpace(username)
    if username == "" {
        return nil, ErrEmptyUsername
    }

    // 2. Llamada al repositorio
    user, err := uc.repo.GetContactByName(username)
    if err != nil {
        return nil, err
    }

    // 3. Usuario no encontrado
    if user == nil {
        return nil, ErrUserNotFound
    }

    return user, nil
}
