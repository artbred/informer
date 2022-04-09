package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Client struct {
	BaseURL string
}

const (
	SendTelegramMessageEndpoint string = "/telegram/send-message"
	CallEndpoint string = "/call"
)

type (
	SendTelegramMessageRequest struct {
		Token string `json:"chat_token"`
		Message string `json:"message"`
	}

	CallRequest struct {
		Phone string `json:"phone"`
		Message string `json:"message"`
	}

	JSONResponse struct {
		Code int `json:"code"`
		Message string `json:"message"`
	}
)

func (c *Client) SendTelegramMessage(message, chatToken string) (err error) {
	url := c.BaseURL + SendTelegramMessageEndpoint

	req := SendTelegramMessageRequest{
		Token:   chatToken,
		Message: message,
	}

	b, err := json.Marshal(req); if err != nil {
		logrus.WithError(err).Errorf("Can't send message to telegram chat")
		return
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(b)); if err != nil {
		logrus.WithError(err).Errorf("Can't send message to telegram chat")
		return
	}

	if res.StatusCode == http.StatusCreated {
		return nil
	}

	return errors.New("unsuccessfully")
}

func (c *Client) Call(message, phone string) (err error) {
	url := c.BaseURL + CallEndpoint

	req := CallRequest{
		Phone:   phone,
		Message: message,
	}

	b, err := json.Marshal(req); if err != nil {
		logrus.WithError(err).Errorf("Can't send message to telegram chat")
		return
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(b)); if err != nil {
		logrus.WithError(err).Errorf("Can't send message to telegram chat")
		return
	}

	if res.StatusCode == http.StatusCreated {
		return nil
	}

	return errors.New("unsuccessfully")
}

func New(baseUrl string) Client {
	return Client{
		BaseURL:  baseUrl,
	}
}