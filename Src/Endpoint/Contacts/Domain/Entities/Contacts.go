package entities

type Contact struct {
    ID        int    `json:"id"`
    UserID    int    `json:"userId"`
    ContactID int    `json:"contactId"`
    CreatedAt string `json:"createdAt"`
}
