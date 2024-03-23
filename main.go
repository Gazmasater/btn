package main

import (
	"fmt"
	"log"
	"time"

	"uslugio.com/USLUGIO/authdata"
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

	// Получение логина и пароля из пакета login
	user := authdata.GetUserCredentials()

	// Ввод логина
	if err := webdriver_utils.InputText(wd, "input[name='email']", user.Email, "логина"); err != nil {
		log.Fatalf("%s", err)
	}

	// Ввод пароля
	if err := webdriver_utils.InputText(wd, "input[name='pass']", user.Password, "пароля"); err != nil {
		log.Fatalf("%s", err)
	}

	// Нажатие на кнопку входа
	if err := webdriver_utils.ClickButton(wd, "button[type='submit']", "входа"); err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("Успешно вошли в систему")

	// Нажатие на кнопку Up
	if err := webdriver_utils.ClickButton(wd, ".up_date-76722", "Up"); err != nil {
		log.Fatalf("%s", err)
	}
}
