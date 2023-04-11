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
	authorService := service.NewAuthorService(dao)
	bookService := service.NewBookService(dao)

	// handler
	handler := handler.NewMicroService(authorService, bookService)

	// grouping router
	authorGroup := app.Group("/author")
	authorRouter(authorGroup, handler)

	// grouping authentication
	authenticationGroup := app.Group("/auth")
	authenticationRouter(authenticationGroup, handler)
}
