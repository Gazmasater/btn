package authdata

// Credentials содержит информацию о логине и пароле
type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
