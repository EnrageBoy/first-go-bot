package main

import (
	"flag"
	"log"

	"./clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	t := mustToken()

	tgClient = telegram.New(tgBotHost, mustToken())

	//fetcher = fetcher.New()

	//processor = processor.New(tgClient)

	//consumer.Start(feacher, processor)
}

func mustToken() string {

	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
}
