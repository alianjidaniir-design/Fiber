package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.Accepts("text/html")
		c.Accepts("json", "text")
		c.Accepts("application/json")
		c.Accepts("text/plain", "application/json")
		c.Accepts("image/png")
		c.Accepts("png")
		return c.SendString("Hello, World!")
	})
	log.Fatal(app.Listen(":3000"))

}
