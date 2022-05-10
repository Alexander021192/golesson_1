package commands

import (
	"fmt"
	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	outputList := c.GetListData(1)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputList)

	// msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Next page", c.GetListData(2)),
	// 	),
	// )

	c.bot.Send(msg)
}

func (c *Commander) NextPage (inputMsg *tgbotapi.Message, currentPage *int) {
	outputList := c.GetListData(*currentPage + 1)
	*currentPage = *currentPage + 1
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputList)
	c.bot.Send(msg)
}

func (c *Commander) GetListData(page int) string {
	products := c.productService.List(page)
	listData := fmt.Sprintf("Products (page %d):\n", page)
	for _, el := range products {
		listData += el.Title + "\n"
	}
	return listData
}
