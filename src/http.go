package src

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/tucnak/telebot.v2"
	"net/http"
)


type SendMessageTelegramRequest struct {
	Token string `json:"chat_token"`
	Message string `json:"message"`
}

func SendTelegramMessage (c echo.Context) error {
	req := &SendMessageTelegramRequest{}
	if err := c.Bind(req); err != nil {
		return JsonResponse(c, http.StatusBadRequest, err.Error())
	}

	chatID, err := GetChatByToken(req.Token); if err != nil {
		return JsonResponse(c, http.StatusBadRequest, err.Error())
	}

	if _, err := Bot.Send(&telebot.Chat{ID:chatID}, req.Message, telebot.NoPreview); err != nil {
		return JsonResponse(c, http.StatusInternalServerError, err.Error())
	}

	return JsonResponse(c, http.StatusCreated, "Send!")
}

type CallPhoneRequest struct {
	Phone string `json:"phone"`
	Message string `json:"message"`
}

func CallPhone (c echo.Context) error {
	req := &CallPhoneRequest{}
	if err := c.Bind(req); err != nil {
		return JsonResponse(c, http.StatusBadRequest, err.Error())
	}

	Call(req.Phone, req.Message)
	return JsonResponse(c, http.StatusCreated, "Send!")
}

func StartHttpServer (serverPort string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/telegram/send-message", SendTelegramMessage)
	e.POST("/call", CallPhone)

	e.Logger.Fatal(e.Start(":"+serverPort))
}
