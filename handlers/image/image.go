package image

import (
	cryptControllers "Github/sarthakpranesh/silvershare/controllers/crypt"
	keyControllers "Github/sarthakpranesh/silvershare/controllers/key"
	"Github/sarthakpranesh/silvershare/controllers/responses"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func EncryptImage(c *fiber.Ctx) error {
	rawUID := c.Get("Authorization")
	rawUID = strings.Replace(rawUID, "Bearer", "", -1)
	rawUID = strings.TrimSpace(rawUID)
	img, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	keyId, err := strconv.ParseUint(c.FormValue("keyId"), 0, 0)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	buf := bytes.NewBuffer(nil)
	file, err := img.Open()
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	_, err = io.Copy(buf, file)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	imgBytes := buf.Bytes()
	key, err := keyControllers.GetKey(uint(keyId), rawUID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	crypted, err := cryptControllers.EncryptAES(imgBytes, key.Secret)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	err = cryptControllers.WriteFile(img.Filename, crypted)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	return c.Download("./temp/" + img.Filename)
}

func DecryptImage(c *fiber.Ctx) error {
	rawUID := c.Get("Authorization")
	rawUID = strings.Replace(rawUID, "Bearer", "", -1)
	rawUID = strings.TrimSpace(rawUID)
	img, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	keyId, err := strconv.ParseUint(c.FormValue("keyId"), 0, 0)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	buf := bytes.NewBuffer(nil)
	file, err := img.Open()
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	_, err = io.Copy(buf, file)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	imgBytes := buf.Bytes()
	key, err := keyControllers.GetKey(uint(keyId), rawUID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
	}
	opImage, err := cryptControllers.DecryptAES(imgBytes, key.Secret)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	err = cryptControllers.WriteFile(img.Filename, opImage)
	if err != nil {
		fmt.Println(err)
		return c.JSON(responses.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
	}
	return c.Download("./temp/" + img.Filename)
}
