package application

import (
    "errors"
    entities "chat/Src/Endpoint/Contacts/Domain/Entities"
    repository "chat/Src/Endpoint/Contacts/Domain/Repository"
)

type GetAllContactsUseCase struct {
    repository repository.IContacts
}

func NewGetAllContactsUseCase(repository repository.IContacts) *GetAllContactsUseCase {
    return &GetAllContactsUseCase{
        repository: repository,
    }
}

func (uc *GetAllContactsUseCase) Execute(userID int) ([]entities.ContactResponse, error) {
    if userID <= 0 {
        return nil, errors.New("userID inválido")
    }

    contacts, err := uc.repository.GetAll(userID)
    if err != nil {
        return nil, err
    }

    return contacts, nil
}
