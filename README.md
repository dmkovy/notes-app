# Сервис заметок
Веб-сервис Go для создания и чтения самоуничтожающихся заметок.

Самый быстрый способ запустить это в вашей собственной среде — использовать Docker Compose:
```
version: "3.9"

services:
  db:
    image: redis:latest
    ports:
      - "6379:6379"
  web:
    build: .
    environment:
      - PORT=3000
      - REDIS_URL=redis://:@db:6379/1
    ports:
      - "3000:3000"
    depends_on:
      - db
```
Команда для запуска в терминале:
```
docker compose up
```