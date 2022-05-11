package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - help\n"+
		"/list - list Products\n"+
		"/next_page - next_page list Products\n"+
		"/get - get\n"+
		"/new - new product\n"+
		"/delete - delete product\n")
	c.bot.Send(msg)
}
