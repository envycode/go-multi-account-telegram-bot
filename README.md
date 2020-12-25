# Multi Account Telegram Bot

POC to prove we can register, listen and remove telegram bot on the fly.

## Setup
require: golang 1.15

```
make dep
make run
```

## Sample Request

### Add new bot

```
curl --location --request POST 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "your-bot-name",
    "token": "your-telegram-token'
```

### Remove bot

```
curl --location --request DELETE 'http://localhost:8080/deregister' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "your-bot-name"
}'
```
