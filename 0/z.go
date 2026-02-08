package main

import "github.com/gofiber/fiber/v3"
import "github.com/gofiber/fiber/v3/extractors"

type User struct {
	ID   int    `params:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := fiber.New()
	var user User
	app.Get("/user/:id", func(c fiber.Ctx) error {
		if err := c.Bind().Body(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(user)
	})

	apikeyExtractor := extractors.Chain(
		extractors.FromHeader("X-Api-Key"),
		extractors.FromQuery("api-key"),
		extractors.FromCookie("api-key"),
	)
	app.Use(func(c fiber.Ctx) error {
		apiKey, err := apikeyExtractor.Extract(c)
		if err != nil {
			return c.Status(401).SendString("API key required")
		}
		return c.Next()
	})

	app.Listen(":3000")

}
