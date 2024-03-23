// Package webdriverutils предоставляет утилиты для работы с веб-драйвером.
package webdriver_utils

import (
	"fmt"

	"github.com/tebeka/selenium"
)

// InputText вводит текст в указанное поле ввода.
func InputText(wd selenium.WebDriver, selector, text, actionDescription string) error {
	inputField, err := wd.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		return fmt.Errorf("ошибка при поиске поля для ввода %s: %s", actionDescription, err)
	}
	if err := inputField.SendKeys(text); err != nil {
		return fmt.Errorf("ошибка при вводе %s: %s", actionDescription, err)
	}
	return nil
}

// ClickButton нажимает на указанную кнопку.
func ClickButton(wd selenium.WebDriver, selector, actionDescription string) error {
	button, err := wd.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		return fmt.Errorf("ошибка при поиске кнопки %s: %s", actionDescription, err)
	}
	if err := button.Click(); err != nil {
		return fmt.Errorf("ошибка при нажатии кнопки %s: %s", actionDescription, err)
	}
	fmt.Printf("Кнопка %s успешно нажата\n", actionDescription)
	return nil
}
