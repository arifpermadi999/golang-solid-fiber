package dto

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func SetResponseSuccess(c *fiber.Ctx, message string) error {
	var res Response
	res.Message = message
	res.Success = true
	res.Status = http.StatusOK
	return c.Status(http.StatusOK).JSON(res)
}
func SetResponseDataSuccess(c *fiber.Ctx, message string, data interface{}) error {
	var res Response
	res.Success = true
	res.Message = message
	res.Status = http.StatusOK
	res.Data = data
	return c.Status(http.StatusOK).JSON(res)
}

func SetResponseError(c *fiber.Ctx, errorMes string, errorLog ...string) error {
	var res Response
	// Set errorLog if provided
	if errorLog != nil {
		println(errorLog[0])
		//set error log
	}
	res.Success = false
	res.Message = errorMes
	res.Status = getStatusError(&errorMes)
	return c.Status(http.StatusOK).JSON(res)
}
func getStatusError(errorMes *string) int {
	switch *errorMes {
	case "record not found":
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
