package main

import (
	"Github/sarthakpranesh/silvershare/handlers/image"
	"Github/sarthakpranesh/silvershare/handlers/key"
	"Github/sarthakpranesh/silvershare/handlers/user"
	"Github/sarthakpranesh/silvershare/middleware"
	"image"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Protected routes, see the middleware
	app.Use(middleware.Authentication)
	app.Use("/user/:name", user.Register)
	app.Use("/key/new", key.NewKey)
	app.Use("/key/all", key.AllKeys)
	app.Use("/key/:id", key.KeyDetails)
	app.Use("/key/share", key.KeyShare)
	app.Use("/img/enc", image.EncryptImage)
	app.Use("/img/dec", image.DecryptImage)

	// Not use Default "*" catcher, handles any route not implemented on server
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendString("Not yet implemented!")
	})

	// app listening to port and other options
	app.Listen("0.0.0.0:8080")
}
