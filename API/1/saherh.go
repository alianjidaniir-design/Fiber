package main

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func compress(a any) any {
	switch a.(type) {
	case int:
		return a.(int)*a.(int) + 1
	case float64:
		return a
	case string:
		return strings.ToLower(a.(string))
	default:
		return a
	}

}

func fib(c fiber.Ctx) error {
	f1 := c.Query("f1")
	g := compress(f1)
	return c.JSON(fiber.Map{"data": g})
}
func main() {
	app := fiber.New()
	app.Get("/compress", fib)
	app.Listen(":3001")
}
