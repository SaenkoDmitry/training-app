# Training Telegram Bot

## Overview
A Telegram bot for tracking workout training sessions.

## Interface

### Preview and workout session
<p align="left">
  <img width="200" height="450" src="/screenshots/preview.png">  
  <img width="200" height="450" src="/screenshots/list_of_workouts+statistics.png">  
  <img width="200" height="450" src="/screenshots/workout_session.png">
</p>

### Export to Excel
<p align="left">
  <img width="200" height="450" src="/screenshots/excel_0.png">
  <img width="200" height="450" src="/screenshots/excel_1.png">
</p>

## Project Structure
- `cmd/main.go` - Main entry point
- `internal/models/` - Database models (User, WorkoutDay, Exercise, Set, WorkoutSession)
- `internal/repository/` - Database repositories for each model
- `internal/service/` - Bot service and message handlers
- `internal/utils/` - Utility functions

## Dependencies
- Go 1.24
- `github.com/go-telegram-bot-api/telegram-bot-api/v5` - Telegram Bot API
- `github.com/pressly/goose/v3` - for database migrations
- `gorm.io/gorm, gorm.io/driver/postgres` - ORM with Postgres driver

## Configuration and secrets
The bot requires:
1. Environment variable `TELEGRAM_TOKEN` containing your Telegram Bot API token (get from @BotFather on Telegram)
2. Environment variable `DATABASE_URL` containing DSN for connection to your database

## Running
```bash
go build -o training-tg-bot ./cmd/main.go
./training-tg-bot
```

## Database
Uses Postgres database. Auto-migrates on startup via github.com/pressly/goose/v3.

## DATABASE_URL for local startup
postgresql://postgres:postgres@localhost/training-bot
