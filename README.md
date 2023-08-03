
# Проект создания множества почтовых ящиков на MailSlurp с помощью API 
Этот проект содержит код для создания сразу нескольких почтовых ящиков с помощью API сервиса MailSlurp.

***
### Требования
Для запуска проекта вам понадобятся следующие компоненты:

- Go 1.20
- Аккаунт и API-ключ с сайта [Ссылка на сайт](https://www.mailslurp.com)

***

1. Установите Go (версия 1.20 или выше) - [Ссылка на загрузку](https://go.dev/doc/install).

2. Склонируйте репозиторий на вашу локальную машину:

   ```bash
   git clone https://github.com/Jhnvlglmlbrt/MailSlurp-CreatingEmail

3. Перейдите в директорию проекта:

   ```bash
   cd MailSlurp-CreatingEmail

4. Установите зависимости проекта:
    ```bash
    go get
    
5. Скопируйте API-Ключ - [Ссылка](https://app.mailslurp.com/dashboard/)

6. Замените "Your API-key" на ваш реальный API-ключ в функции createClient.

    ```bash
    apiKey := "Your API-key"

7. Запустите проект: 

    ```bash
    go run main.go

