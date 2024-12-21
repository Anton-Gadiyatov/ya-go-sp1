# Calculator Web Service

Этот проект представляет собой веб-сервис, который позволяет пользователям отправлять арифметические выражения через HTTP POST-запрос и получать результат вычисления. Сервис поддерживает базовые арифметические операции, такие как сложение, вычитание, умножение и деление, а также обработку скобок.

---

## Основные возможности

1. **Эндпоинт:**

   - `/api/v1/calculate` — единственный эндпоинт для обработки запросов.

2. **Метод запроса:**

   - `POST`

3. **Формат запроса:**

   - Тело запроса должно быть в формате JSON:
     ```json
     {
       "expression": "выражение, которое ввёл пользователь"
     }
     ```

4. **Формат ответа:**
   - Успешный ответ (код 200):
     ```json
     {
       "result": "результат выражения"
     }
     ```
   - Ошибка (код 422, если выражение не валидно):
     ```json
     {
       "error": "Expression is not valid"
     }
     ```
   - Ошибка (код 500, если произошла внутренняя ошибка):
     ```json
     {
       "error": "Internal server error"
     }
     ```

---

## Требования

- **Go 1.16 или выше**
- Убедитесь, что у вас установлен Go и настроена переменная окружения `GOPATH`.

---

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/anton-gadiyatov/ya-go-sp1-calc.git
cd ya-go-sp1-calc
```

### 2. Инициализация модуля Go

```bash
go mod init ya-go-sp1-calc
```

Затем установите зависимости:

```bash
go mod tidy
```

### 3. Запуск сервера

Для запуска сервера выполните команду:

```bash
go run main.go
```

По умолчанию сервер будет доступен по адресу http://localhost:8080.

## Пример использования

Отправка запроса
Используйте curl для отправки запроса:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
 -H "Content-Type: application/json" \
 -d '{"expression": "1 + 2 * (3+4/2-(1+2))*2+1"}'
```

Пример успешного ответа

```json
{
  "result": 10
}
```

Пример ответа с ошибкой
Если выражение содержит недопустимые символы:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
 -H "Content-Type: application/json" \
 -d '{"expression": "1 + 2 \* a"}'
```

Ответ:

```json
{
  "error": "Expression is not valid"
}
```

## Тестирование

Для тестирования сервиса можно использовать следующие примеры запросов:

### Успешный запрос:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
 -H "Content-Type: application/json" \
 -d '{"expression": "2 + 2"}'
```

Ответ:

```json
{
  "result": 4
}
```

### Невалидное выражение:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
 -H "Content-Type: application/json" \
 -d '{"expression": "1 + \* 2"}'
```

Ответ:

```json
{
  "error": "Expression is not valid"
}
```

### Внутренняя ошибка:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
 -H "Content-Type: application/json" \
 -d '{"expression": "1 / 0"}'
```

Ответ:

```json
{
  "error": "Internal server error"
}
```
