package repository

import entities "chat/Src/Endpoint/Contacts/Domain/Entities"

type IContacts interface {
	GetAll(userID int) ([]entities.Contact, error)
	CreateContact(contact entities.Contact) error
	DeleteContact(userID int, contactID int) error
	Exists(userID int, contactID int) (bool, error)
}
