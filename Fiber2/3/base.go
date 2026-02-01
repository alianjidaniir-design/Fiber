package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	micro := fiber.New()
	app.Static("/", "./rep")

	app.Mount("/mount", micro)
	micro.Get("/doe", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(":3000"))
}
