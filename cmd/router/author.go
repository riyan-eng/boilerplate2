package router

import (
	"boilerplate/cmd/handler"

	"github.com/gofiber/fiber/v2"
)

func authorRouter(router fiber.Router, handler *handler.MicroServiceServer) {
	// router.Get("/")
	router.Post("/", handler.CreateAuthor)
}
