package handler

import (
	"boilerplate/cmd/dto"
	handlerReqres "boilerplate/cmd/handler/reqres"
	serviceReqres "boilerplate/cmd/service/reqres"
	"boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func (m *MicroServiceServer) RegisterAdmin(c *fiber.Ctx) error {
	payload := new(handlerReqres.RegisterAdminRequest)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
			Data:    err.Error(),
			Message: pkg.MESSAGE_BAD_REQUEST,
		})
	}
	if err := pkg.Validate(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
			Data:    err,
			Message: pkg.MESSAGE_BAD_REQUEST,
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
		if serviceRes.Error.(*pq.Error).Code.Name() == "unique_violation" {
			return c.Status(fiber.StatusConflict).JSON(pkg.ResponseJson{
				Data:    pkg.PqErrGenerate(serviceRes.Error),
				Message: pkg.MESSAGE_BAD_REQUEST,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ResponseJson{
			Data:    serviceRes.Error.Error(),
			Message: pkg.MESSAGE_BAD_SYSTEM,
		})
	}
	return c.JSON(pkg.ResponseJson{
		Message: pkg.MESSAGE_OK_CREATE,
	})
}
