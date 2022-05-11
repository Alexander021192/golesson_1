package commands

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Delete(inputMsg *tgbotapi.Message) {
	productTitleDel := inputMsg.CommandArguments()
	if productTitleDel != "" {
		file, err := os.Open("products.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var newProducts string
		for scanner.Scan() {
			fmt.Println(scanner.Text(), productTitleDel)
			if scanner.Text() != productTitleDel {
				newProducts += scanner.Text() + "\n"
			}
		}

		// Write new products
		err = ioutil.WriteFile("products.txt", []byte(newProducts), 0644)
		if err != nil {
			log.Fatal(err)
		}
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("success delete: %s", productTitleDel))
		c.bot.Send(msg)
	}
}
