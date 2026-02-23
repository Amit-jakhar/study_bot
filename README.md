# Go Study Bot

Telegram bot that sends scheduled Go study topic reminders and lets you mark each topic as done.

## What It Does
- Sends reminders at fixed times (10:30, 22:00, 23:00 IST).
- Each reminder includes the current topic.
- When you send `Done`, it moves to the next topic.

## Setup
1. Create a Telegram bot and get your token.
2. Find your chat ID.
3. Fill `.env`:

```env
TELEGRAM_BOT_TOKEN=your_token_here
TELEGRAM_CHAT_ID=your_chat_id_here
```

## Run
```bash
go run .
```

## Notes
- Timezone is forced to `Asia/Kolkata` (IST) regardless of server location.
- If `Asia/Kolkata` fails to load, it falls back to `UTC`.
- Topics are defined in `main.go`.
- `.env` is ignored by git via `.gitignore`.
