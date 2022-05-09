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

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("failed get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("successfully parsed args: %v,\nproduct title: %v\n", idx, product.Title),
	)
	c.bot.Send(msg)
}
