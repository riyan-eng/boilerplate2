package handler

import (
	"boilerplate/cmd/dto"
	handlerReqres "boilerplate/cmd/handler/reqres"
	serviceReqres "boilerplate/cmd/service/reqres"
	"boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
)

func (m *MicroServiceServer) CreateAuthor(c *fiber.Ctx) error {
	payload := new(handlerReqres.CreateAuthorRequest)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
			Data:    err.Error(),
			Message: "bad",
		})
	}
	serviceReq := serviceReqres.CreateAuthorRequest{
		Context: c.Context(),
		Item: dto.Author{
			Name:        payload.Name,
			Address:     payload.Address,
			PhoneNumber: payload.PhoneNumber,
		},
	}
	serviceRes := m.authorService.CreateAuthor(&serviceReq)
	if serviceRes.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(pkg.ResponseJson{
			Data:    serviceRes.Error.Error(),
			Message: "bad",
		})
	}
	return c.JSON(pkg.ResponseJson{
		Data:    "successfully insert data",
		Message: "ok",
	})
}
