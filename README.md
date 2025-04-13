Проект URL Shortener API

URL Shortener API — это простой сервис для сокращения длинных ссылок, который включает в себя регистрацию и вход пользователей, защиту маршрутов через JWT и автоматическую очистку «протухших» ссылок.
🚀 Что умеет

    Регистрация

        POST /register

        Тело JSON:

    {
      "username": "myuser",
      "password": "mypassword"
    }

    Регистрирует нового пользователя.

Вход (Login)

    POST /login

    Тело JSON:

{
  "username": "myuser",
  "password": "mypassword"
}

В ответе вернёт JWT-токен:

    { "token": "<JWT_TOKEN>" }

Создание короткой ссылки

    POST /shortener

    Заголовок:

Authorization: Bearer <JWT_TOKEN>

Тело JSON:

{
  "original_url": "https://example.com/very/long/path",
  "ttl_seconds": 3600
}

В ответе вернёт короткий код:

    { "short_code": "abc123XYZ" }

Список своих ссылок

    GET /shortener

    Тот же заголовок Authorization

    Вернёт массив объектов:

    [
      {
        "id": 1,
        "short_code": "abc123XYZ",
        "original_url": "https://example.com/very/long/path",
        "created_at": "2025-04-14T12:34:56Z",
        "ttl_second": 3600
      }
    ]

Редирект по короткому коду

    GET /shortener/{code}

    Перенаправит вас (HTTP 303) на исходный URL.

Информация о себе

    GET /me

    Заголовок Authorization

    Вернёт ваш user_id:

        { "user_id": 1 }

🛠 Как запустить
1. С помощью Docker Compose

    Клонируйте репозиторий и перейдите в папку:

git clone https://github.com/yourusername/URL-Shortener-API.git
cd URL-Shortener-API

Создайте файл .env рядом с docker-compose.yml:

DB_HOST=url_shortener_db
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=url_shortener
API_PORT=8080
JWT_SECRET=your-super-secret-key

Запустите всё одной командой:

    docker-compose up --build

    Готово!

        БД PostgreSQL доступна на localhost:5432

        API слушает http://localhost:8080

2. Локально без Docker

    Установите Go 1.18+ и PostgreSQL

    Создайте .env с теми же переменными

    Установите зависимости и запустите:

    go mod tidy
    go run cmd/main.go

    Откройте http://localhost:8080

🔧 Технологии

    Go — язык

    gorilla/mux — маршрутизация

    PostgreSQL — база данных

    robfig/cron — удаление просроченных ссылок

    bcrypt — безопасное хеширование паролей

    JWT — авторизация пользователей

✨ Что можно добавить

    Веб‑интерфейс для удобства

    Swagger‑документацию

    Улучшенная валидация входных данных

    Rate limiting и защита от брутфорса

    Unit и интеграционные тесты