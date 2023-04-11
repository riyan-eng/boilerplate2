package router

import (
	"boilerplate/cmd/handler"

	"github.com/gofiber/fiber/v2"
)

func authenticationRouter(router fiber.Router, handler *handler.MicroServiceServer) {
	router.Post("/register_admin", handler.RegisterAdmin)
}
