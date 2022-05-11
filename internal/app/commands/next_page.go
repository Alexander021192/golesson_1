package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) NextPage(inputMsg *tgbotapi.Message) {
	currentPage, err := c.productService.CurrentPage()
	if err != nil {
		log.Printf("problem with get current page")
	}
	*currentPage += 1
	outputList := c.GetListData(*currentPage)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputList)
	c.bot.Send(msg)
}
