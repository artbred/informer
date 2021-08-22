package src

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomToken (length int) string {
	b := make([]rune, length)

	rand.Seed(time.Now().UnixNano())

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func JsonResponse (c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]interface{}{
		"Code": code,
		"Message": message,
	})
}
