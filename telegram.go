package main

import (
	"log"
	"net/url"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

func initTelegramBot() *telebot.Bot {
	b, err := telebot.NewBot(telebot.Settings{
		Token:     conf.TelegramAPIKey,
		Poller:    &telebot.LongPoller{Timeout: PollerTimeout * time.Second},
		Verbose:   false,
		ParseMode: "html",
	})

	if err != nil {
		log.Fatal(err)
	}

	return b
}

func logTelegram(message string) {
	message = "anote-robot:" + getCallerInfo() + message
	rec := &telebot.Chat{
		ID: int64(conf.LogAccount),
	}
	bot.Send(rec, message)
}

func notificationTelegram(message string) {
	rec := &telebot.Chat{
		ID: int64(conf.LogAccount),
	}
	bot.Send(rec, message)
}

func logTelegramService(message string) error {
	m, _ := url.QueryUnescape(message)
	message, _ = url.PathUnescape(m)
	var err error

	rec := &telebot.Chat{
		ID: int64(conf.LogAccount),
	}

	if strings.Contains(message, "no data for this key") {
		_, err = bot.Send(rec, message, telebot.NoPreview, telebot.Silent)
	} else {
		_, err = bot.Send(rec, message)
	}
	return err
}