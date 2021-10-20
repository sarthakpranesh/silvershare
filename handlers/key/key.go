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
	rawUserId := c.Get("Authorization")
	rawUserId = strings.Replace(rawUserId, "Bearer", "", -1)
	rawUserId = strings.TrimSpace(rawUserId)
	user_id, err := strconv.ParseUint(rawUserId, 0, 0)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	k := keyControllers.Key{
		Secret: key,
		UserId: uint(user_id),
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
	rawUserId := c.Get("Authorization")
	rawUserId = strings.Replace(rawUserId, "Bearer", "", -1)
	rawUserId = strings.TrimSpace(rawUserId)
	user_id, err := strconv.ParseUint(rawUserId, 0, 0)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	keys, err := keyControllers.AllUserKeys(uint(user_id))
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
	rawUserId := c.Get("Authorization")
	rawUserId = strings.Replace(rawUserId, "Bearer", "", -1)
	rawUserId = strings.TrimSpace(rawUserId)
	user_id, err := strconv.ParseUint(rawUserId, 0, 0)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	id, err := strconv.ParseUint(c.Params("id"), 0, 0)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	key, err := keyControllers.GetKey(uint(id), uint(user_id))
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
	rawUserId := c.Get("Authorization")
	rawUserId = strings.Replace(rawUserId, "Bearer", "", -1)
	rawUserId = strings.TrimSpace(rawUserId)
	user_id, err := strconv.ParseUint(rawUserId, 0, 0)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	var shared keyControllers.Shared
	err = c.BodyParser(&shared)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: "Invalid Request Body",
			Status:  400,
		})
	}
	err = keyControllers.ShareKey(uint(user_id), shared)
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
