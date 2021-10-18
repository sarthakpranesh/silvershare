package user

import (
	"Github/sarthakpranesh/silvershare/controllers"
	"Github/sarthakpranesh/silvershare/controllers/responses"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user controllers.User
	err := c.BodyParser(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: "Invalid Request Body",
			Status:  400,
		})
	}
	err = controllers.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	return c.JSON(user)
}
