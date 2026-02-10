package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.RouteChain("/test").Get(func(c fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})
	app.Get("/users:id", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"route": c.FullPath(),
		})

	})

	log.Fatal(app.Listen(":3013"))
}
