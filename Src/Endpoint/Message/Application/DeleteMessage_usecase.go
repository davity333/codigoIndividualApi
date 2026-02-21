package application

import repository "chat/Src/Endpoint/Message/Domain/Repository"

type DeleteMessageUseCase struct {
	repository repository.IMessage
}


func NewDeleteMessageUseCase(repository repository.IMessage) *DeleteMessageUseCase {
	return &DeleteMessageUseCase{
		repository: repository,
	}
}

func (uc *DeleteMessageUseCase) Execute(messageId int) error {
    err := uc.repository.DeleteMessage(messageId)
    if err != nil {
        return err
    }
    return nil
}
