package src

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/ratelimit"
	"gopkg.in/tucnak/telebot.v2"
	"net/http"
)

var limiter ratelimit.Limiter

type SendMessageTelegramRequest struct {
	Token string `json:"chat_token"`
	Message string `json:"message"`
}

func SendTelegramMessage (c echo.Context) error {
	limiter.Take()

	req := &SendMessageTelegramRequest{}
	if err := c.Bind(req); err != nil {
		return JsonResponse(c, http.StatusBadRequest, err.Error())
	}

	chatID, err := GetChatByToken(req.Token); if err != nil {
		return JsonResponse(c, http.StatusBadRequest, err.Error())
	}

	if _, err := Bot.Send(&telebot.Chat{ID:chatID}, req.Message); err != nil {
		return JsonResponse(c, http.StatusInternalServerError, err.Error())
	}

	return JsonResponse(c, http.StatusCreated, "Send!")
}

func StartHttpServer (serverPort string) {
	limiter = ratelimit.New(1)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/telegram/send-message", SendTelegramMessage)
	e.Logger.Fatal(e.Start(serverPort))
}
