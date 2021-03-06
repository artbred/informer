package src

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

type Chat struct {
	ChatToken string `db:"chat_token"`
	ChatID int64 `db:"chat_id"`
}

func AddChat (token string, chatID int64) (err error) {
	_, err = db.Exec("INSERT INTO chats (chat_token, chat_id) VALUES ($1, $2) ON CONFLICT (chat_id) DO UPDATE SET chat_token=$1", token, chatID)

	if err != nil {
		logrus.WithError(err).Errorf("Can't add chat %s with token %d to database", chatID, token)
	}

	return
}

func GetChat (chatID int64) (chat Chat, exists bool, err error) {
	row := db.QueryRow("SELECT * FROM chats WHERE chat_id=$1", chatID)
	exists = true

	if err = row.Scan(&chat.ChatToken, &chat.ChatID); err != nil {
		if err == sql.ErrNoRows {
			return Chat{}, false, nil
		}
		logrus.WithError(err).Errorf("Can't get chat %d", chatID)
	}

	return
}

func GetChatByToken (token string) (chatID int64, err error) {
	row := db.QueryRow("SELECT chat_id FROM chats WHERE chat_token=$1", token)

	if err = row.Scan(&chatID); err != nil {
		if err == sql.ErrNoRows {
			return chatID, errors.New("Chat not found")
		}
		logrus.WithError(err).Errorf("Can't scan value in order to get chat by token %s", token)
		return chatID, err
	}

	return
}

func init () {
	database, err := sql.Open("sqlite3", "./data/informer.db"); if err != nil {
		go panic(err)
	}

	if err := database.Ping(); err != nil {
		go panic(err)
	}

	db = database
}