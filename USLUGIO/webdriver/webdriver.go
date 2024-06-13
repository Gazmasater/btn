package webdriver

import (
	"fmt"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	seleniumPath = "/usr/local/bin/chromedriver"
	port         = 4444
)

// StartWebDriver запускает Selenium WebDriver и создает удаленный драйвер
func StartWebDriver() (selenium.WebDriver, error) {
	opts := []selenium.ServiceOption{}
	_, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запуске Chrome WebDriver: %s", err)
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			// Устанавливаем режим headless
			"--headless",
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании удаленного драйвера: %s", err)
	}

	return wd, nil
}
