package main

import (
	"log"
	"os"

	"github.com/Alexander021192/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil { // If we got a message

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}
	
		}
	}
}

func listCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message, productSevice *product.Service) {
	products := productSevice.List()
	outputList := "Here all products:\n\n"
	for _, el := range products {
		outputList += el.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputList)
	bot.Send(msg)
}

func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - help\n"+
		"/list - list Products\n")
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: " +inputMsg.Text)
	msg.ReplyToMessageID = inputMsg.MessageID
	bot.Send(msg)
}