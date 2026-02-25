package application

import (
    "errors"
    entities "chat/Src/Endpoint/Contacts/Domain/Entities"
    repository "chat/Src/Endpoint/Contacts/Domain/Repository"
)

type CreateContactUseCase struct {
    repository repository.IContacts
}

func NewCreateContactUseCase(repository repository.IContacts) *CreateContactUseCase {
    return &CreateContactUseCase{
        repository: repository,
    }
}

func (uc *CreateContactUseCase) Execute(contact entities.Contact) error {

    // Validación de IDs
    if contact.UserID <= 0 || contact.ContactID <= 0 {
        return errors.New("IDs inválidos")
    }

    // No agregarse a sí mismo
    if contact.UserID == contact.ContactID {
        return errors.New("no puedes agregarte a ti mismo como contacto")
    }

    // Validar duplicado
    exists, err := uc.repository.Exists(contact.UserID, contact.ContactID)
    if err != nil {
        return err
    }
    if exists {
        return errors.New("el contacto ya está agregado")
    }

    // Crear contacto
    return uc.repository.CreateContact(contact)
}
