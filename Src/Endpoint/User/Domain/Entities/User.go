package entities

type User struct {
	ID       int `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Role string `json:"rol"`
}

func NewUser(id int, username, email, password, firstName, lastName, role string) *User {
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
		FirstName: firstName,
		LastName: lastName,
		Role: role,
	}
}