package controller

import (
	"github.com/gofiber/fiber/v2"
	"koriebruh/management/dto"
	"koriebruh/management/service"
	"net/http"
)

type ItemController interface {
	CreateItem(ctx *fiber.Ctx) error
	FindAllByCategory(ctx *fiber.Ctx) error
}

type ItemControllerImpl struct {
	service.ItemService
}

func NewItemController(itemService service.ItemService) *ItemControllerImpl {
	return &ItemControllerImpl{ItemService: itemService}
}

func (controller ItemControllerImpl) CreateItem(ctx *fiber.Ctx) error {

	var request dto.ItemRequest
	token := ctx.Cookies("token")

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	err := controller.ItemService.Create(ctx.Context(), token, request)
	if err != nil { // <-- if got error in service or repo
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data: map[string]string{
			"message": "success created new item",
		},
	})

}

func (controller ItemControllerImpl) FindAllByItem(ctx *fiber.Ctx) error {

	item, err := controller.ItemService.FindAllItem(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   item,
	})

}