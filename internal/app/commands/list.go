package commands

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	page, err := strconv.Atoi(args)
	if err != nil {
		page = 1
	}
	outputList := c.GetListData(page)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputList)

	// msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Next page", c.GetListData(2)),
	// 	),
	// )

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
