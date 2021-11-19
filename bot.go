package main

import (
	"os"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		panic(err)
	}
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		//checking type of message to be text
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				//greeting message
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a bot checking for unique items that needed for recipes(in plans also for prophecies) in game Path of Exile(RU)")
				bot.Send(msg)
			default:
				item := update.Message.Text
				//searching for recipes
				checkresult := SearchRecipes(item)
				if len(checkresult) == 0 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No matches was found")
					bot.Send(msg)
				}
				//sending all matches
				for _, elem := range checkresult {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, elem)
					bot.Send(msg)
				}
			}
		} else {
			//asking to type something
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter item's name")
			bot.Send(msg)
		}
	}
}
