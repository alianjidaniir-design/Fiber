package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()
	app.Get("/", func(req fiber.Req, res fiber.Res) error {
		return res.SendString("Hello from Express_style handler")
	})
	app.Use(func(req fiber.Req, res fiber.Res, next func() error) error {
		if req.IP() == "192.168.0.2" {
			return res.SendStatus(fiber.StatusForbidden)
		}
		return next()
	})
	app.Use(func(req fiber.Req, res fiber.Res, next func()) {
		if req.Get("X-Skip") == "true" {
			return // stop the chain without calling next
		}
		next()
	})

	log.Fatal(app.Listen(":3000"))

}
