package main

import (
	"boilerplate/cmd/router"
	"boilerplate/config"
	"boilerplate/migration"
	"log"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	file, err := os.OpenFile("./log/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "${time} ${status} - ${method} ${path}\n",
		TimeFormat: "2006/01/02 15:04:05",
	}))
	log.SetOutput(file)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	router.NewRouter(app)
	app.Listen(":3000")
}
