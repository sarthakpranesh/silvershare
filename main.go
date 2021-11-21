package main

import (
	"Github/sarthakpranesh/silvershare/connections"
	"Github/sarthakpranesh/silvershare/handlers/image"
	"Github/sarthakpranesh/silvershare/handlers/key"
	"Github/sarthakpranesh/silvershare/handlers/user"
	"Github/sarthakpranesh/silvershare/middleware"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// creating required file directories to process images
	err := os.Mkdir("temp", 0777)
	if err != nil {
		fmt.Println("Unable to create `temp` directory for image processing, error: ", err)
	}

	// initiate the app database
	connections.PostgresConnector()

	app := fiber.New()

	// Protected routes, see the middleware
	app.Use(cors.New())
	app.Use(middleware.Authentication)
	app.Post("/user", user.Register)
	app.Get("/key/new", key.NewKey)
	app.Get("/key/all", key.AllKeys)
	app.Get("/key/:id", key.KeyDetails)
	app.Post("/key/share", key.KeyShare)
	app.Post("/img/enc", image.EncryptImage)
	app.Post("/img/dec", image.DecryptImage)

	// Not use Default "*" catcher, handles any route not implemented on server
	app.Use("*", func(c *fiber.Ctx) error {
		return c.SendString("Not yet implemented!")
	})

	// app listening to port and other options
	app.Listen("0.0.0.0:" + os.Getenv("PORT"))
}
