package handlers

import (
	"fmt"

	"github.com/arifpermadi999/core-echo-golang/constants"
	"github.com/arifpermadi999/core-echo-golang/dto"
	"github.com/arifpermadi999/core-echo-golang/usecases"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthUsecase *usecases.AuthUsecase
}

func ProviderAuthHandler(au *usecases.AuthUsecase) AuthHandler {
	return AuthHandler{AuthUsecase: au}
}

func (au *AuthHandler) Login(c *fiber.Ctx) error {
	authDto := *new(dto.AuthDto)
	err := c.BodyParser(&authDto)
	fmt.Print(authDto.Password)
	if err != nil {
		return dto.SetResponseError(c, err.Error())
	}
	data, err := au.AuthUsecase.Login(&authDto)
	if err != nil {
		return dto.SetResponseError(c, err.Error())
	}
	return dto.SetResponseDataSuccess(c, "login berhasil", data)
}

func (au *AuthHandler) Register(c *fiber.Ctx) error {
	register := *new(dto.RegisterUser)
	err := c.BodyParser(&register)

	if err != nil {
		return dto.SetResponseError(c, constants.ErrBindingParamStr, err.Error())
	}
	data, err := au.AuthUsecase.Register(&register)
	if err != nil {
		return dto.SetResponseError(c, err.Error())
	}
	return dto.SetResponseDataSuccess(c, "register success", data)
}
