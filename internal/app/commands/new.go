package commands

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) New(inputMsg *tgbotapi.Message) {
	productTitle := inputMsg.CommandArguments()
	if productTitle != "" {
		//Append products
		file, err := os.OpenFile("products.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		if _, err := file.WriteString(productTitle + "\n"); err != nil {
			log.Fatal(err)
		}

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("success created new product: %s", productTitle))
		c.bot.Send(msg)
	}
}
