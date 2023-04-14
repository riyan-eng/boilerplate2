package handler

import (
	serviceReqres "boilerplate/cmd/service/reqres"
	"boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
)

func (m *MicroServiceServer) Logout(c *fiber.Ctx) error {
	serviceReq := serviceReqres.LogoutRequest{
		Context: c.Context(),
		UserID:  c.Locals("userID").(string),
	}
	serviceRes := m.authenticationService.Logout(&serviceReq)
	if serviceRes.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ResponseJson{
			Data:    serviceRes.Error.Error(),
			Message: "bad",
		})
	}
	return c.JSON(pkg.ResponseJson{
		Data:    "successfully logout",
		Message: "ok",
	})
}
