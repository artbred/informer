package main

import (
	"github.com/artbred/informer/src"
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

func main () {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	go src.StartHttpServer(os.Getenv("SERVER_PORT"))
	b := src.Bot

	b.Handle("/start@" + b.Me.Username, func(m *tb.Message) {
		token := src.RandomToken(8)

		if err := src.AddChat(token, m.Chat.ID); err != nil {
			b.Send(m.Chat, "Please try again later")
			return
		}

		b.Send(m.Chat, token)
	})

	err := b.SetCommands([]tb.Command{{Text: "start", Description: "Activate chat and get token"}})
	if err != nil {
		panic(err)
	}

	b.Start()
}