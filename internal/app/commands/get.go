package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args: ", args)
		return
	}

	product := c.productService.ReadProducts(idx, idx+1)
	productTitle := ""
	lenList, err := c.productService.Count()
	if err != nil {
		return
	}
	if idx >= lenList || idx < 0 {
		productTitle += fmt.Sprintf("not found product with number %d \n", idx)
	}
	for _, el := range product {
		productTitle += el.Title + "\n"
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("successfully parsed args: %v,\nproduct title: %v\n", idx, productTitle),
	)
	c.bot.Send(msg)
}
