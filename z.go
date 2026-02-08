package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Listen("app.souk", fiber.ListenConfig{
		ListenerNetwork:    fiber.NetworkUnix,
		UnixSocketFileMode: 0777,
	})

}
