package main

import (
	"flag"
	"log"

	"github.com/DEHbNO4b/tgbot-advertisement/internal/clients/telegram"
)

const (
	tgBotHost = "api.telegram.bot"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	//fetcher = fetcher.New(tgClient)

	//processor = processor.New(tgClient)

	//consumer.Start(fetcher,processor)
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for acces to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
