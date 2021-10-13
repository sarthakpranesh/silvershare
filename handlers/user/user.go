package user

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	var name string = c.Params("name")
	return c.SendString(name)
}
