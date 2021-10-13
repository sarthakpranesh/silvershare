package image

import "github.com/gofiber/fiber/v2"

func EncryptImage(c *fiber.Ctx) error {
	return c.SendString("Enc Image")
}

func DecryptImage(c *fiber.Ctx) error {
	return c.SendString("Dec Image")
}
