package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	if !fiber.IsChild() {
		fmt.Println("fiber is not child")
	} else {
		fmt.Println("fiber is child")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")

}
