package fileio

import (
	"uslugio.com/USLUGIO/buttonstorage"

	"encoding/json"
	"fmt"
	"os"
)

func ReadButtonsFromFile(filePath string) ([]buttonstorage.Button, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла: %s", err)
	}
	defer file.Close()

	var buttons []buttonstorage.Button
	if err := json.NewDecoder(file).Decode(&buttons); err != nil {
		return nil, fmt.Errorf("ошибка при декодировании файла JSON: %s", err)
	}
	return buttons, nil
}
