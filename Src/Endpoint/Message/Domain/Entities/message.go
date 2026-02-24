package entities

type Message struct {
	ID       int    `json:"idmessage"`
	Content  string `json:"content"`
	ReceiveId int    `json:"receiverId"`
	SenderId   int    `json:"senderId"`
	TimeMessage int64 `json:"timeMessage"`
}
