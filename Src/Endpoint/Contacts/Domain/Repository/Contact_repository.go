package repository

import entities "chat/Src/Endpoint/Contacts/Domain/Entities"
import userEntities "chat/Src/Endpoint/User/Domain/Entities"

type IContacts interface {
	GetAll(userID int) ([]entities.Contact, error)
	CreateContact(contact entities.Contact) error
	GetContactByName(contactName string) (*userEntities.User, error)
	DeleteContact(userID int, contactID int) error
	Exists(userID int, contactID int) (bool, error)
}
