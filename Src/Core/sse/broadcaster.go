package sse

import (
	"sync"

	"github.com/gorilla/websocket"
)

type MessageEvent struct {
	ID         int64  `json:"id"`
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
}

type Broadcaster struct {
	clients map[int]map[*websocket.Conn]bool
	mu      sync.RWMutex
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		clients: make(map[int]map[*websocket.Conn]bool),
	}
}

// Subscribe adds a client to receive messages
func (b *Broadcaster) Subscribe(userID int, conn *websocket.Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.clients[userID] == nil {
		b.clients[userID] = make(map[*websocket.Conn]bool)
	}
	b.clients[userID][conn] = true
}

// Unsubscribe removes a client from receiving messages
func (b *Broadcaster) Unsubscribe(userID int, conn *websocket.Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if connections, ok := b.clients[userID]; ok {
		if _, exists := connections[conn]; exists {
			delete(connections, conn)
			conn.Close()
		}
		if len(connections) == 0 {
			delete(b.clients, userID)
		}
	}
}

// Broadcast sends a message to all subscribers of a specific user
func (b *Broadcaster) Broadcast(userID int, message MessageEvent) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if connections, ok := b.clients[userID]; ok {
		for conn := range connections {
			err := conn.WriteJSON(message)
			if err != nil {
				// We can't safely remove from map while iterating in Go without care,
				// but usually the read loop on the other side handles exact closure.
				conn.Close()
			}
		}
	}
}

// BroadcastToMultiple sends message to multiple users
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
