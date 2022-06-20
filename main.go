package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env, %s", err)
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(telegramToken)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s \n", update.Message.From.UserName, update.Message.Text)

			regex, _ := regexp.Compile(`elah+`)

			isContainElah := regex.MatchString(update.Message.Text)

			var answer string
			if isContainElah {
				answer = fmt.Sprintf("elah mulu lu %s", update.Message.From.UserName)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
			// else {
			// 	answer = "aman"
			// }

			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
			// msg.ReplyToMessageID = update.Message.MessageID

			// bot.Send(msg)
		}
	}
}
