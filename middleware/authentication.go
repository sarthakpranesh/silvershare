package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	var uid string = c.Get("Authorization")
	if len(uid) == 0 {
		return c.SendString("Not Authorized")
	}
	return c.Next()
}
