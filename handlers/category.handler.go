package handlers

import (
	"strconv"

	"github.com/arifpermadi999/core-echo-golang/constants"
	"github.com/arifpermadi999/core-echo-golang/dto"
	"github.com/arifpermadi999/core-echo-golang/models"
	"github.com/arifpermadi999/core-echo-golang/usecases"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	CategoryUsecase *usecases.CategoryUsecase
}

func ProviderCategoryHandler(u *usecases.CategoryUsecase) CategoryHandler {
	return CategoryHandler{CategoryUsecase: u}
}

func (h *CategoryHandler) GetAllCategory(c *fiber.Ctx) error {
	data := h.CategoryUsecase.GetAllDataCategory()
	return dto.SetResponseDataSuccess(c, constants.ReceiveStr, data)
}

func (h *CategoryHandler) ViewCategory(c *fiber.Ctx) error {
	id, errb := strconv.Atoi(c.Params("id"))
	if errb != nil {
		return dto.SetResponseError(c, constants.ErrConvertStr, errb.Error())
	}
	model := new(models.Category)
	model.Id = id

	data, error := h.CategoryUsecase.ViewDataCategory(model)
	if error != nil {
		return dto.SetResponseError(c, error.Error())
	}
	return dto.SetResponseDataSuccess(c, constants.ReceiveStr, data)
}

func (h *CategoryHandler) SaveCategory(c *fiber.Ctx) error {
	model := new(models.Category)
	errb := c.BodyParser(model)
	if errb != nil {
		return dto.SetResponseError(c, constants.ErrBindingParamStr, errb.Error())
	}
	data, error := h.CategoryUsecase.SaveCategory(model)
	if error != nil {
		return dto.SetResponseError(c, constants.ErrSaveStr, error.Error())
	}
	return dto.SetResponseDataSuccess(c, constants.SaveStr, data)
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	model := new(models.Category)
	errb := c.BodyParser(model)
	if errb != nil {
		return dto.SetResponseError(c, constants.ErrBindingParamStr, errb.Error())
	}
	data, error := h.CategoryUsecase.UpdateCategory(model)
	if error != nil {
		return dto.SetResponseError(c, constants.ErrUpdateStr, error.Error())
	}
	return dto.SetResponseDataSuccess(c, constants.UpdateStr, data)
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, errb := strconv.Atoi(c.Params("id"))
	if errb != nil {
		return dto.SetResponseError(c, constants.ErrConvertStr)
	}
	model := new(models.Category)
	model.Id = id

	error := h.CategoryUsecase.DeleteCategory(model)
	if error != nil {
		return dto.SetResponseError(c, constants.ErrDeleteStr, error.Error())
	}
	return dto.SetResponseError(c, constants.DeleteStr)
}
