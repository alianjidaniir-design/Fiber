package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Bind(fiber.Map{
			"Title": "Hello World",
		})
		return c.SendString("emam")
	})
	app.Get("/Tavan", func(c *fiber.Ctx) error {
		return c.Render("xxx.tmpl", fiber.Map{})
	})
	app.Listen(":3002")

}
