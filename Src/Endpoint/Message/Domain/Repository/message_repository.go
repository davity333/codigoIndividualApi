package repository

import entities "chat/Src/Endpoint/Message/Domain/Entities"

type IMessage interface {
	GetMessagesByUserId(userId int) ([]*entities.Message, error)
	SendMessage(message *entities.Message) error
	DeleteMessage(messageId int) error
}