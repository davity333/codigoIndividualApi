package application

import (
	entities "chat/Src/Endpoint/Message/Domain/Entities"
	repository "chat/Src/Endpoint/Message/Domain/Repository"
)

type SendMessageUseCase struct {
	repository repository.IMessage
}

func NewSendMessageUseCase(repository repository.IMessage) *SendMessageUseCase {
	return &SendMessageUseCase{
		repository: repository,
	}
}

func (uc *SendMessageUseCase) Execute(message *entities.Message) error {
    return (uc.repository).SendMessage(message)
}
