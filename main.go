package main

import (
	"log"

	tgClient "Telegram_bot/clients/telegram"
	event_consumer "Telegram_bot/consumer/event-consumer"
	"Telegram_bot/events/telegram"
	"Telegram_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, "6771436601:AAHxgTSyt8utNdwcQUjNgSdFMqsbNZRlSDk"),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stoped", err)
	}
}

// func mustToken() string {

// 	token := flag.String(
// 		"tg-bot-token",
// 		"",
// 		"token for access to telegram bot",
// 	)

// 	flag.Parse()

// 	if *token == "" {
// 		log.Fatal("token is not specified")
// 	}

// 	return *token
// }
