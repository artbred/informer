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

		chat, exists, err := src.GetChat(m.Chat.ID); if err != nil {
			b.Send(m.Chat, "Please try again later")
			return
		}

		if exists {
			b.Send(m.Chat, chat.ChatToken)
			return
		}

		if err := src.AddChat(token, m.Chat.ID); err != nil {
			b.Send(m.Chat, "Please try again later")
			return
		}

		b.Send(m.Chat, token)
	})

	b.Handle("/call", func(m *tb.Message) {
		src.Call(m.Payload)
	})

	err := b.SetCommands([]tb.Command{
		{Text: "start", Description: "Activate chat and get token"},
		{Text: "call", Description: "Call to phone and say something"},
	})

	if err != nil {
		panic(err)
	}

	b.Start()
}