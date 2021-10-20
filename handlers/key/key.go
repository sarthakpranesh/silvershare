package key

import (
	keyControllers "Github/sarthakpranesh/silvershare/controllers/key"
	"Github/sarthakpranesh/silvershare/controllers/responses"
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewKey(c *fiber.Ctx) error {
	rawUID := c.Get("Authorization")
	rawUID = strings.Replace(rawUID, "Bearer", "", -1)
	rawUID = strings.TrimSpace(rawUID)
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	k := keyControllers.Key{
		Secret: key,
		UserId: rawUID,
	}
	err = keyControllers.CreateKey(&k)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	return c.JSON(k)
}

func AllKeys(c *fiber.Ctx) error {
	rawUID := c.Get("Authorization")
	rawUID = strings.Replace(rawUID, "Bearer", "", -1)
	rawUID = strings.TrimSpace(rawUID)
	keys, err := keyControllers.AllUserKeys(rawUID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	return c.JSON(keys)
}

func KeyDetails(c *fiber.Ctx) error {
	rawUID := c.Get("Authorization")
	rawUID = strings.Replace(rawUID, "Bearer", "", -1)
	rawUID = strings.TrimSpace(rawUID)
	id, err := strconv.ParseUint(c.Params("id"), 0, 0)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	key, err := keyControllers.GetKey(uint(id), rawUID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	return c.JSON(key)
}

func KeyShare(c *fiber.Ctx) error {
	rawUID := c.Get("Authorization")
	rawUID = strings.Replace(rawUID, "Bearer", "", -1)
	rawUID = strings.TrimSpace(rawUID)
	var shared keyControllers.Shared
	err := c.BodyParser(&shared)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: "Invalid Request Body",
			Status:  400,
		})
	}
	err = keyControllers.ShareKey(rawUID, shared)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	return c.JSON(struct {
		Message string `json:"message"`
		Status  uint   `json:"status"`
	}{
		Message: "Shared",
		Status:  200,
	})
}
