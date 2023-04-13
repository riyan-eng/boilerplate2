package handler

import (
	"boilerplate/cmd/dto"
	handlerReqres "boilerplate/cmd/handler/reqres"
	serviceReqres "boilerplate/cmd/service/reqres"
	"boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
)

func (m *MicroServiceServer) RegisterAdmin(c *fiber.Ctx) error {
	payload := new(handlerReqres.RegisterAdminRequest)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
			Data:    err.Error(),
			Message: "bad",
		})
	}
	if err := pkg.Validate(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
			Data:    err,
			Message: "bad",
		})
	}
	serviceReq := serviceReqres.RegisterAdminRequest{
		Context: c.Context(),
		Item: dto.RegisterAdmin{
			UserName:    payload.UserName,
			Password:    payload.Password,
			Email:       payload.Email,
			PhoneNumber: payload.PhoneNumber,
		},
	}
	serviceRes := m.authenticationService.RegisterAdmin(&serviceReq)
	if serviceRes.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ResponseJson{
			Data:    serviceRes.Error.Error(),
			Message: "bad",
		})
	}
	return c.JSON(pkg.ResponseJson{
		Data:    "successfully insert data",
		Message: "ok",
	})
}
