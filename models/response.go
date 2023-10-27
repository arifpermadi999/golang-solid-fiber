package models

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponseSuccess(c *fiber.Ctx, message string) error {
	var res Response
	res.Message = message
	res.Status = http.StatusOK
	return c.Status(http.StatusOK).JSON(res)
}
func SetResponseDataSuccess(c *fiber.Ctx, message string, data interface{}) error {
	var res Response
	res.Message = message
	res.Status = http.StatusOK
	res.Data = data
	return c.Status(http.StatusOK).JSON(res)
}

func SetResponseError(c *fiber.Ctx, errorMes string) error {
	var res Response
	res.Message = errorMes
	res.Status = http.StatusInternalServerError

	return c.Status(http.StatusOK).JSON(res)
}
