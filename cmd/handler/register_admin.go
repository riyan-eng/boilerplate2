package handler

import (
	handlerReqres "boilerplate/cmd/handler/reqres"
	"boilerplate/pkg"
	"fmt"

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
	fmt.Println(payload)
	return c.JSON(pkg.ResponseJson{
		Data:    "successfully insert data",
		Message: "ok",
	})
}
