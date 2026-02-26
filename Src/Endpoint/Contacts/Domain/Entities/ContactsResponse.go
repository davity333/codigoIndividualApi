package entities

type ContactResponse struct {
    ID         int    `json:"id"`
    UserID     int    `json:"userId"`
    ContactID  int    `json:"contactId"`
    CreatedAt  string `json:"createdAt"`
    FirstName  string `json:"firstName"`
    LastName   string `json:"lastName"`
    Username   string `json:"username"`
    Email      string `json:"email"`
}
