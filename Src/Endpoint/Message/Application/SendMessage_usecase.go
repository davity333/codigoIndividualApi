package application

import (
	"time"

	"chat/Src/Core/sse"
	entities "chat/Src/Endpoint/Message/Domain/Entities"
	repository "chat/Src/Endpoint/Message/Domain/Repository"
)

type SendMessageUseCase struct {
	repository  repository.IMessage
	broadcaster *sse.Broadcaster
}

func NewSendMessageUseCase(repository repository.IMessage) *SendMessageUseCase {
	return &SendMessageUseCase{
		repository:  repository,
		broadcaster: nil, // Will be set by DI
	}
}

// SetBroadcaster sets the broadcaster for real-time notifications
func (uc *SendMessageUseCase) SetBroadcaster(broadcaster *sse.Broadcaster) {
	uc.broadcaster = broadcaster
}

func (uc *SendMessageUseCase) Execute(message *entities.Message) error {
	err := (uc.repository).SendMessage(message)
	if err != nil {
		return err
	}

	// Broadcast message to receiver in real-time if broadcaster is set
	if uc.broadcaster != nil {
		sseEvent := sse.MessageEvent{
			ID:         int64(message.ID),
			SenderID:   message.SenderId,
			ReceiverID: message.ReceiveId,
			Content:    message.Content,
			CreatedAt:  time.Now().Format(time.RFC3339),
		}
		// Send to receiver for real-time notification
		uc.broadcaster.Broadcast(message.ReceiveId, sseEvent)
	}

	return nil
}
