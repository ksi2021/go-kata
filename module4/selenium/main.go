package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"time"
)

const maxTries = 5

func main() {
	// прописываем конфигурацию для драйвера
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// добавляем в конфигурацию драйвера настройки для chrome
	chrCaps := chrome.Capabilities{
		W3C: true,
	}
	caps.AddChrome(chrCaps)

	// переменная нашего веб драйвера
	var wd selenium.WebDriver
	var err error
	// прописываем адрес нашего драйвера
	urlPrefix := selenium.DefaultURLPrefix
	// немного костылей чтобы драйвер не падал
	i := 1
	for i < maxTries {
		wd, err = selenium.NewRemote(caps, urlPrefix)
		if err != nil {
			log.Println(err)
			i++
			continue
		}
		break
	}
	// после окончания программы завершаем работу драйвера
	defer wd.Quit()

	// сразу обращаемся к странице с поиском вакансии по запросу
	page := 1         // номер страницы
	query := "golang" // запрос
	wd.Get(fmt.Sprintf("https://career.habr.com/vacancies?page=%d&q=%s&type=all", page, query))

	data, err := wd.FindElement(selenium.ByCSSSelector, ".search-total")
	vacancy, err := data.Text()

	fmt.Println(vacancy)

	// ждем 60 секунд, чтобы успеть посмотреть результат
	time.Sleep(60 * time.Second)
}
