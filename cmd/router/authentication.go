package router

import (
	"boilerplate/cmd/handler"
	"boilerplate/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func authenticationRouter(router fiber.Router, handler *handler.MicroServiceServer) {
	router.Post("/login", handler.Login)
	router.Post("/refresh", handler.RefreshToken)
	router.Post("/logout", middleware.AuthorizeJwt(), handler.Logout)
	router.Post("/register_admin", middleware.AuthorizeJwt(), handler.RegisterAdmin)
}
