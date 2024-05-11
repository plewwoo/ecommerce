package common

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserAuthorization struct {
	IsSuccess bool
	Code      string
	Message   string
}

func getLastPathComponent(path string) string {
	components := strings.Split(path, "/")
	lastComponent := components[len(components)-1]
	if lastComponent == "favicon.ico" {
		return ""
	}
	return lastComponent
}

func Logout(c *fiber.Ctx) error {

	token := c.Get("Authorization")

	if token == "" {
		return errors.New("no token")
	}

	return nil
}
