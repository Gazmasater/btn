package authdata

// Credentials содержит информацию о логине и пароле
type Credentials struct {
	Email    string
	Password string
}

// GetUserCredentials возвращает логин и пароль пользователя
func GetUserCredentials() Credentials {
	// В данном примере логин и пароль жестко заданы, но в реальной реализации
	// они могут быть получены из конфигурационного файла, среды окружения и т.д.
	return Credentials{
		Email:    "lew1968@list.ru",
		Password: "Gazmasterpro358",
	}
}
