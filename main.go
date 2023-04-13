package main

import (
	"boilerplate/cmd/router"
	"boilerplate/config"
	"boilerplate/migration"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU / 2)
	}
	config.ConnDatabase()
	migration.Start()
	config.ConnRedis()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	router.NewRouter(app)
	app.Listen(":3000")
}
