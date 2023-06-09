package handler

import (
	"boilerplate/cmd/dto"
	handlerReqres "boilerplate/cmd/handler/reqres"
	serviceReqres "boilerplate/cmd/service/reqres"
	"boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
)

func (m *MicroServiceServer) Login(c *fiber.Ctx) error {
	payload := new(handlerReqres.LoginRequest)
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
	serviceReq := serviceReqres.LoginRequest{
		Context: c.Context(),
		Issuer:  string(c.Request().Host()),
		Item: dto.Login{
			UserName: payload.UserName,
			Password: payload.Password,
		},
	}
	serviceRes := m.authenticationService.Login(&serviceReq)
	if serviceRes.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
			Data:    serviceRes.Error.Error(),
			Message: pkg.MESSAGE_BAD_SYSTEM,
		})
	}
	handleRes := handlerReqres.LoginResponse{
		AccessToken:  serviceRes.AccessToken,
		RefreshToken: serviceRes.RefreshToken,
		ExpiredAt:    serviceRes.ExpiredAt,
		UserName:     serviceRes.Item.UserName,
		UserTypeName: serviceRes.Item.UserTypeName,
		Name:         serviceRes.Item.Name,
		Email:        serviceRes.Item.Email,
		PhoneNumber:  serviceRes.Item.PhoneNumber,
	}
	return c.JSON(pkg.ResponseJson{
		Data:    handleRes,
		Message: "successfully login.",
	})
}
