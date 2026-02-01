package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api2 := app.Group("/api2")

	v1 := api2.Group("/v1", ff)
	v1.Get("/v1", ff)
	v1.Get("/messages", ff)

	v2 := api2.Group("/v2", gg) // /api/v2
	v2.Get("/list", gg)         // /api/v2/list
	v2.Get("/user", gg)

	app.Route("/test", func(api fiber.Router) {
		api.Get("/foo", func(c *fiber.Ctx) error {
			return c.SendString("foo")
		}).Name("foo")
		api.Get("/bar", func(c *fiber.Ctx) error {
			return c.SendString("bar")
		}).Name("bar") // /test/bar (name: test.bar)
	}, "test.")

	log.Fatal(app.Listen(":3000"))

}

func ff(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func gg(c *fiber.Ctx) error {
	return c.SendString("good bye")
}
