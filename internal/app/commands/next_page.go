package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) NextPage(inputMsg *tgbotapi.Message) {
	currentPage := c.productService.CurrentPage()

	*currentPage += 1
	outputList := c.GetListData(*currentPage)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputList)
	c.bot.Send(msg)
}
