# Используйте официальный образ Golang как базовый образ
FROM golang:latest

# Установите рабочую директорию в контейнере
WORKDIR /go/src/app

# Установка необходимых пакетов для установки Google Chrome
RUN apt-get update && apt-get install -y wget

# Копирование .deb файла Google Chrome внутрь контейнера
COPY google-chrome-stable_current_amd64-107.0.5304.63.deb /tmp/

# Установка Google Chrome из .deb файла
RUN apt-get install -y /tmp/google-chrome-stable_current_amd64-107.0.5304.63.deb

# Удаление .deb файла после установки
RUN rm /tmp/google-chrome-stable_current_amd64-107.0.5304.63.deb



# Копирование файла chromedriver в контейнер
COPY chromedriver /usr/local/bin/

# Установка прав на выполнение для chromedriver
RUN chmod +x /usr/local/bin/chromedriver

# Вывод версии chromedriver
RUN chromedriver --version

# Вывод версии Google Chrome
RUN google-chrome --version

# Инициализация модуля Go
RUN go mod init uslugio.com


# Установка Selenium и его драйверов
RUN go get "github.com/tebeka/selenium"
RUN go get "github.com/tebeka/selenium/chrome"



# Скопируйте файлы вашего Go приложения в текущую рабочую директорию в контейнере
COPY . .


# Обновление и оптимизация списка зависимостей Go
RUN go mod tidy


# Соберите приложение внутри контейнера
RUN go install -v ./...

# Укажите команду для запуска вашего приложения при старте контейнера
CMD ["uslugio.com"]
