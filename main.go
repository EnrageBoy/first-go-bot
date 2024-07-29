package main

import (
	"flag"
	"log"
)

func main() {
	//token = flag.Get(token)

	//tgClient = telegram.New(token)

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