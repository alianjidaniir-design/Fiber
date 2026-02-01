package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {

	app := fiber.New()
	// Provide a minimal config

	// Or extend your config for customization
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			c.Status(fiber.StatusUnauthorized)
			c.Set("WWW-Authenticate", "Basic realm=\"Forbidden\"")

			return c.SendString("Access Granted to the Realm!")

		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		// این کد فقط برای کاربرانی اجرا می‌شود که توسط basicauth احراز هویت شده‌اند
		return c.SendString("Welcome to the Protected Root! Authentication Successful.")
	})
	app.Listen(":3001")
}
