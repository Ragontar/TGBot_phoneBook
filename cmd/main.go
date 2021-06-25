package main

import (
	"fmt"
	"github.com/Ragontar/anusRipper/pkg/setup"
)

func main() {
	token, err := setup.GetTelegramAPItoken()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Got Telegram API token: %s", token)
}
