package main

import (
	"log"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

)

func main() {
	// Load .env if present (local dev convenience)
	_ = godotenv.Load()

	// 1. Token aur Chat ID env se lo
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("missing TELEGRAM_BOT_TOKEN env var")
	}

	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	if chatIDStr == "" {
		log.Fatal("missing TELEGRAM_CHAT_ID env var")
	}

	myChatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("invalid TELEGRAM_CHAT_ID: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Bot authorized on account %s", bot.Self.UserName)

	// 2. Updated Topics List (2-3 ghante wale topics)
	topics := []string{
		"Basics & Variables",
		"Flow Control (If/For)",
		"Functions & Blank Identifier",
		"Slices (Dynamic Arrays)",
		"Maps (Key-Value)",
		"Structs (Custom Types)",
		"Pointers & Memory",
		"Methods & Receivers",
		"Interfaces (Abstraction)",
		"Error Handling (The Go Way)",
		"Goroutines (Concurrency)",
		"Channels (Communication)",
		"JSON Handling",
		"Standard Library Basics",
	}

	currentIndex := 0

	// 3. Fixed Time Reminder Logic (Subah 10:30, Raat 10:00 aur 11:00)
	// go func() {
	// 	// India ka Location load karo
	// 	loc, err := time.LoadLocation("Asia/Kolkata")
	// 	if err != nil {
	// 		log.Println("Error loading location:", err)
	// 		// Agar server par location load na ho, toh fallback default UTC par rakho
	// 		loc = time.UTC
	// 	}

	// 	for {
	// 		// DHAYAN SE DEKHO: Yahan .In(loc) add kiya hai. Ab red line gayab!
	// 		currentTime := time.Now().In(loc).Format("15:04")

	// 		if currentTime == "10:30" || currentTime == "22:00" || currentTime == "23:00" {
	// 			if currentIndex < len(topics) {
	// 				text := "Bhai, Go time! 🚀\n\nAbhi ka topic hai: *" + topics[currentIndex] + "*\n\nKhatam karke 'Done' mark karo!"
	// 				msg := tgbotapi.NewMessage(myChatID, text)
	// 				msg.ParseMode = "Markdown"
	// 				bot.Send(msg)
	// 			}
	// 			time.Sleep(90 * time.Second)
	// 		}
	// 		time.Sleep(30 * time.Second)
	// 	}
	// }()

	go func() {
		// India Location (IST) load karo
		loc, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			log.Println("Error loading location, using UTC fallback:", err)
			loc = time.UTC
		}

		for {
			// Local testing ke liye hum specific time nahi, interval use kar rahe hain
			if currentIndex < len(topics) {
				text := "🚀 *Go Study Alert (Test Mode)*\n\nTopic: *" + topics[currentIndex] + "*\n\nBhai padhai shuru kar! Khatam karke 'Done' reply kar."
				msg := tgbotapi.NewMessage(myChatID, text)
				msg.ParseMode = "Markdown"
				bot.Send(msg)
			}

			// TEST: Har 10 second mein reminder
			time.Sleep(10 * time.Second)

			// NOTE: Jab bot live karna ho, toh yahan time check wala logic wapas daal dena
			// currentTime := time.Now().In(loc).Format("15:04")
			// if currentTime == "10:30" || currentTime == "22:00" || currentTime == "23:00" { ... }
		}
	}()

	// 4. Message Handling (Done mark karne ke liye)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "Done" {
			if currentIndex < len(topics) {
				msgText := "Shabaash! *" + topics[currentIndex] + "* completed. ✅\nAgla topic agle scheduled time par milega."
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				msg.ParseMode = "Markdown"
				bot.Send(msg)

				currentIndex++ // Agle topic par move karega
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bhai tune saare topics khatam kar diye! Party? 🥳")
				bot.Send(msg)
			}
		}
	}
}
