package main

import "github.com/gofiber/fiber/v3"

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

	app.Get("/car", func(c fiber.Ctx) error {
		return c.Redirect().To("/err")
	})
	app.Get("/err", func(c fiber.Ctx) error {
		return c.SendString("Hi fiber")
	})
	app.Listen(":3000")

}
