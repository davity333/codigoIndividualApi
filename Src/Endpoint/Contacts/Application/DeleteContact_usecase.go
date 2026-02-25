package application

import (
    "errors"
    repository "chat/Src/Endpoint/Contacts/Domain/Repository"
)

type DeleteContactUseCase struct {
    repository repository.IContacts
}

func NewDeleteContactUseCase(repository repository.IContacts) *DeleteContactUseCase {
    return &DeleteContactUseCase{
        repository: repository,
    }
}

func (uc *DeleteContactUseCase) Execute(userID int, contactID int) error {
    if userID <= 0 || contactID <= 0 {
        return errors.New("IDs inválidos")
    }

    if userID == contactID {
        return errors.New("no puedes eliminarte a ti mismo como contacto")
    }

    return uc.repository.DeleteContact(userID, contactID)
}
