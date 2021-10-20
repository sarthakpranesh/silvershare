package user

import (
	"Github/sarthakpranesh/silvershare/controllers/responses"
	controllers "Github/sarthakpranesh/silvershare/controllers/user"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	rawUID := c.Get("Authorization")
	rawUID = strings.Replace(rawUID, "Bearer", "", -1)
	rawUID = strings.TrimSpace(rawUID)
	var user controllers.User
	err := c.BodyParser(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: "Invalid Request Body",
			Status:  400,
		})
	}
	user.Uid = rawUID
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
