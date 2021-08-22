package src

import (
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"time"
)

var Bot *tb.Bot

func Poller() tb.Poller {
	return tb.Poller(&tb.LongPoller{Timeout: 15 * time.Second})
}

func init () {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: Poller(),
	})

	if err != nil {
		go panic(err)
	}

	Bot = b
}

