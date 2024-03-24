package fileio

import (
	"encoding/json"
	"fmt"
	"os"

	"uslugio.com/USLUGIO/authdata"
)

func ReadCredentialsFromFile(filePath string) ([]authdata.Credentials, error) {
	// Открытие файла
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла: %s", err)
	}
	defer file.Close()

	// Создание декодера JSON
	decoder := json.NewDecoder(file)

	// Чтение и декодирование JSON
	var credentials []authdata.Credentials
	if err := decoder.Decode(&credentials); err != nil {
		return nil, fmt.Errorf("ошибка при декодировании файла JSON: %s", err)
	}

	return credentials, nil
}
