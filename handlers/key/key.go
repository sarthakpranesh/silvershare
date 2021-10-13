package key

import "github.com/gofiber/fiber/v2"

func NewKey(c *fiber.Ctx) error {
	return c.SendString("New key")
}

func AllKeys(c *fiber.Ctx) error {
	return c.SendString("All Keys")
}

func KeyDetails(c *fiber.Ctx) error {
	return c.SendString("Key Details")
}

func KeyShare(c *fiber.Ctx) error {
	return c.SendString("Sharing key")
}
