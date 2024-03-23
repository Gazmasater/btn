package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	// Путь к ChromeDriver
	const (
		seleniumPath = "/home/master/chromedriver_linux64/chromedriver"
		port         = 4444
	)

	// Запуск Selenium WebDriver
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if err != nil {
		log.Fatalf("Ошибка при запуске Chrome WebDriver: %s", err)
	}
	defer service.Stop()

	// Настройка ChromeOptions
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
	}
	caps.AddChrome(chromeCaps)

	// Создание удаленного драйвера
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Ошибка при создании удаленного драйвера: %s", err)
	}
	defer wd.Quit()

	// Открытие страницы для входа
	if err := wd.Get("https://uslugio.com/users/login"); err != nil {
		log.Fatalf("Ошибка при открытии страницы входа: %s", err)
	}

	// Ждем, пока страница загрузится
	time.Sleep(1 * time.Second)

	// Поиск поля для ввода логина и ввод логина
	loginInput, err := wd.FindElement(selenium.ByCSSSelector, "input[name='email']")
	if err != nil {
		log.Fatalf("Ошибка при поиске поля для ввода логина: %s", err)
	}
	if err := loginInput.SendKeys("lew1968@list.ru"); err != nil {
		log.Fatalf("Ошибка при вводе логина: %s", err)
	}

	// Поиск поля для ввода пароля и ввод пароля
	passwordInput, err := wd.FindElement(selenium.ByCSSSelector, "input[name='pass']")
	if err != nil {
		log.Fatalf("Ошибка при поиске поля для ввода пароля: %s", err)
	}
	if err := passwordInput.SendKeys("Gazmasterpro358"); err != nil {
		log.Fatalf("Ошибка при вводе пароля: %s", err)
	}

	// Ждем, чтобы пользователь успел увидеть ввод
	time.Sleep(1 * time.Second)

	// Поиск кнопки входа и нажатие на нее
	loginButton, err := wd.FindElement(selenium.ByCSSSelector, "button[type='submit']")
	if err != nil {
		log.Fatalf("Ошибка при поиске кнопки входа: %s", err)
	}
	if err := loginButton.Click(); err != nil {
		log.Fatalf("Ошибка при нажатии кнопки входа: %s", err)
	}

	fmt.Println("Успешно вошли в систему")

	// Поиск кнопки Up по классу
	upButton, err := wd.FindElement(selenium.ByCSSSelector, ".up_date-76722")
	if err != nil {
		log.Fatalf("Ошибка при поиске кнопки Up: %s", err)
	}

	// Проверка на успешность нахождения кнопки
	if upButton != nil {
		fmt.Println("Кнопка Up найдена")

		// Нажатие на кнопку Up
		if err := upButton.Click(); err != nil {
			log.Fatalf("Ошибка при нажатии кнопки Up: %s", err)
		}

		fmt.Println("Кнопка Up успешно нажата")
	} else {
		fmt.Println("Кнопка Up не найдена")
	}
}
