package sse

import (
	"sync"
)

type MessageEvent struct {
	ID         int64  `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
}

type ClientConn struct {
	UserID   int
	Messages chan MessageEvent
}

type Broadcaster struct {
	clients map[int]map[chan MessageEvent]bool
	mu      sync.RWMutex
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		clients: make(map[int]map[chan MessageEvent]bool),
	}
}

// Subscribe adds a client to receive messages
func (b *Broadcaster) Subscribe(userID int) chan MessageEvent {
	b.mu.Lock()
	defer b.mu.Unlock()

	messages := make(chan MessageEvent, 10)

	if b.clients[userID] == nil {
		b.clients[userID] = make(map[chan MessageEvent]bool)
	}
	b.clients[userID][messages] = true

	return messages
}

// Unsubscribe removes a client from receiving messages
func (b *Broadcaster) Unsubscribe(userID int, ch chan MessageEvent) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if channels, ok := b.clients[userID]; ok {
		if _, exists := channels[ch]; exists {
			delete(channels, ch)
			close(ch)
		}
		if len(channels) == 0 {
			delete(b.clients, userID)
		}
	}
}

// Broadcast sends a message to all subscribers of a specific user
func (b *Broadcaster) Broadcast(userID int, message MessageEvent) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if channels, ok := b.clients[userID]; ok {
		for ch := range channels {
			select {
			case ch <- message:
			default:
				// Channel full, skip to avoid blocking
			}
		}
	}
}

// BroadcastToMultiple sends message to multiple users (e.g., conversation participants)
func (b *Broadcaster) BroadcastToMultiple(userIDs []int, message MessageEvent) {
	for _, userID := range userIDs {
		b.Broadcast(userID, message)
	}
}

// GetConnectedUsers returns count of connected users
func (b *Broadcaster) GetConnectedUsers() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.clients)
}
