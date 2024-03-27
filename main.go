package main

import (
	"fmt"
	"log"
	"time"

	"uslugio.com/USLUGIO/buttonstorage"
	"uslugio.com/USLUGIO/fileio"
	"uslugio.com/USLUGIO/webdriver"
	"uslugio.com/USLUGIO/webdriver_utils"
)

func main() {
	// Путь к ChromeDriver
	wd, err := webdriver.StartWebDriver()
	if err != nil {
		log.Fatalf("ошибка при запуске WebDriver: %s", err)
	}
	defer wd.Quit()

	// Открытие страницы для входа
	if err := wd.Get("https://uslugio.com/users/login"); err != nil {
		log.Fatalf("Ошибка при открытии страницы входа: %s", err)
	}

	// Ждем, пока страница загрузится
	time.Sleep(1 * time.Second)

	filePath := "USLUGIO/authdata/config.json"

	credentials, err := fileio.ReadCredentialsFromFile(filePath)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}

	for _, cred := range credentials {
		fmt.Printf("Логин: %s, Пароль: %s\n", cred.Login, cred.Password)
	}

	for _, cred := range credentials {

		// Ввод логина
		if err := webdriver_utils.InputText(wd, "input[name='email']", cred.Login, "логина"); err != nil {
			log.Fatalf("%s", err)
		}

		// Ввод пароля
		if err := webdriver_utils.InputText(wd, "input[name='pass']", cred.Password, "пароля"); err != nil {
			log.Fatalf("%s", err)
		}

		// Нажатие на кнопку входа
		if err := webdriver_utils.ClickButton(wd, "button[type='submit']", "входа"); err != nil {
			log.Fatalf("%s", err)
		}

		// Определение пути к файлу кнопок в зависимости от пары авторизации
		buttonFilePath := fmt.Sprintf("USLUGIO/buttonstorage/uslugio/%s.json", cred.Login)
		println("buttonFilePath", buttonFilePath)
		buttons, err := fileio.ReadButtonsFromFile(buttonFilePath)
		if err != nil {
			log.Fatalf("ошибка при чтении файла кнопок: %s", err)
		}

		for _, button := range buttons {
			time.Sleep(2 * time.Second) // Пауза между нажатиями кнопок

			// Нажатие на кнопку
			err := webdriver_utils.ClickButton(wd, button.Selector, button.Name)
			if err != nil {
				// Если кнопка не найдена, записать в лог и продолжить выполнение цикла
				log.Printf("Ошибка при нажатии кнопки %s: %s", button.Name, err)
				continue
			}
		}

		time.Sleep(1 * time.Second)

		err = webdriver_utils.ClickButton(wd, "a.dropdown-toggle.btn.btn-link", buttonstorage.LoginMap[cred.Login])
		if err != nil {
			log.Fatalf("Ошибка при нажатии кнопки loginMap[cred.Login]: %s", err)
		} else {
			fmt.Println("Кнопка  успешно нажата", buttonstorage.LoginMap[cred.Login])
		}

		time.Sleep(1 * time.Second)

		// Нажатие на кнопку "Выход"
		if err := webdriver_utils.ClickButton(wd, "a[href*='logout']", "Выход"); err != nil {
			log.Fatalf("Ошибка при нажатии кнопки 'Выход': %s", err)
		}

		fmt.Println("Кнопка 'Выход' успешно нажата")
	}

}
