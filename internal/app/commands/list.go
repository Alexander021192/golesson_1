package commands

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	
	currentPage := c.productService.CurrentPage()
	page, err := strconv.Atoi(args)
	if err != nil {
		*currentPage = 1
	} else {
		*currentPage = page
	}
	
	outputList := c.GetListData(*currentPage)
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputList)

	// msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Next page", c.GetListData(2)),
	// 	),
	// )

	c.bot.Send(msg)
}

func (c *Commander) GetListData(page int) string {
	start := 0
	end, _ := c.productService.Count()
	if end == 0 {
		return "empty list product"
	}

	offset := (page - 1) * 5

	last := 1
	if end % 5 == 0 {
		last = 0
	}
	if offset >= end || offset < 0 {
		return fmt.Sprintf("wrong number page %d. Max page is %d", page, end/5 + last)
	}

	if offset < end {
		start = offset
	}
	if offset+5 < end {
		end = offset + 5
	}
	products := c.productService.ReadProducts(start, end)
	listData := fmt.Sprintf("Products (page %d):\n", page)
	for _, el := range products {
		listData += el.Title + "\n"
	}
	return listData
}
