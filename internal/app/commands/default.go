package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)
	msg.ReplyToMessageID = inputMsg.MessageID
	c.bot.Send(msg)
}

func (c *Commander) HandleUpdate(update tgbotapi.Update){
	if update.Message != nil { // If we got a message

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		default:
			c.Default(update.Message)
		}

	}
}