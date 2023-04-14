package handler

import (
	handlerReqres "boilerplate/cmd/handler/reqres"
	serviceReqres "boilerplate/cmd/service/reqres"
	"boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
)

func (m *MicroServiceServer) RefreshToken(c *fiber.Ctx) error {
	payload := new(handlerReqres.RefreshTokenRequest)
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
	serviceReq := serviceReqres.RefreshTokenRequest{
		Context:      c.Context(),
		RefreshToken: payload.RefreshToken,
		Issuer:       string(c.Request().Host()),
	}
	serviceRes := m.authenticationService.RefreshToken(&serviceReq)
	if serviceRes.Error != nil {
		if serviceRes.Error.Error() == pkg.ERROR_REQUEST {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg.ResponseJson{
				Message: pkg.MESSAGE_UNAUTHORIZED,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
			Data:    serviceRes.Error.Error(),
			Message: pkg.MESSAGE_BAD_SYSTEM,
		})
	}
	handlerRes := handlerReqres.RefreshTokenResponse{
		AccessToken:  serviceRes.AccessToken,
		RefreshToken: serviceRes.RefreshToken,
		ExpiredAt:    serviceRes.ExpiredAt,
	}
	return c.JSON(pkg.ResponseJson{
		Data:    handlerRes,
		Message: "successfully refresh.",
	})
}
