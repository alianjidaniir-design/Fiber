package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		c.AllParams()
		c.Accepts("application/json")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": c.Params("name")})

	})

	app.Get("/user/*", func(c *fiber.Ctx) error {
		c.AllParams()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": c.Params("*")})
	})

	app.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})
	log.Fatal(app.Listen(":3003"))

}
