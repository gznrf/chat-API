# chat API

## Описание
Chat API — это API, разработанное для работы в рамках чата, оно выполняет функции работы с пользователями (добавлением и удалением их из базы), паролями (хеш), а так же токен (по нему дается доступ пользователю).

## Функциональные возможности
- **Создание**: Возможность добавления новых пользователей и создания чатов между ними при помощи WebSocket.
- **Чтение**: Получение информации об истории чатов пользователей.
- **Аутентификация**: Поддержка аутентификации пользователей с использованием JWT.

## Технологии
- **Язык программирования**: GOlang
- **Библиотеки**: gorm.io/gorm, spf13/viper, /dgrijalva/jwt-go, gorilla/mux, gorilla/websocket
- **База данных**: PostgreSQL
