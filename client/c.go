package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Use("/user/:id", func(c fiber.Ctx) error {
		fiber.Locals[string](c, "user", "john")
		fiber.Locals[int](c, "age", 25)

		return c.Next()

	})

	app.Get("/user/*", func(c fiber.Ctx) error {
		name := fiber.Locals[string](c, "user")
		age := fiber.Locals[int](c, "age")
		return c.JSON(fiber.Map{"name": name, "age": age})
	})

	app.Listen(":3004")

}
