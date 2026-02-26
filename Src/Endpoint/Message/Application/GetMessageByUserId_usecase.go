package application

import (
	entities "chat/Src/Endpoint/Message/Domain/Entities"
	repository "chat/Src/Endpoint/Message/Domain/Repository"
)

type GetMessageByUserIdUseCase struct {
	repository repository.IMessage
}

func NewGetMessageByUserIdUseCase(repository repository.IMessage) *GetMessageByUserIdUseCase {
	return &GetMessageByUserIdUseCase{
		repository: repository,
	}
}

func (uc *GetMessageByUserIdUseCase) Execute(senderId int, receiveId int) ([]*entities.Message, error) {
	messages, err := uc.repository.GetMessagesByUserId(senderId, receiveId)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
