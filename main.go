package main

import (
	"Github/sarthakpranesh/silvershare/connections"
	"Github/sarthakpranesh/silvershare/handlers/image"
	"Github/sarthakpranesh/silvershare/handlers/key"
	"Github/sarthakpranesh/silvershare/handlers/user"
	"Github/sarthakpranesh/silvershare/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Protected routes, see the middleware
	app.Use(middleware.Authentication)
	app.Post("/user", user.Register)
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

	// initiate the app database
	_, err := connections.PostgresConnector()
	if err != nil {
		os.Exit(1)
	}

	// app listening to port and other options
	app.Listen("localhost:8080")
}
