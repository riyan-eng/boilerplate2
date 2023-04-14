package router

import (
	"boilerplate/cmd/handler"
	"boilerplate/cmd/repository"
	"boilerplate/cmd/service"
	"boilerplate/config"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App) {
	// casbin enforcer
	// enforcer:= config.Casbin()

	// data access object
	dao := repository.NewDao(config.PostgreSQLDB)

	// service
	authenticationService := service.NewAuthenticationService(dao)

	// handler
	handler := handler.NewMicroService(authenticationService)

	// enforce
	enforcer := config.Casbin()

	// grouping router
	authenticationGroup := app.Group("/auth")
	authenticationRouter(authenticationGroup, handler, enforcer)
}
