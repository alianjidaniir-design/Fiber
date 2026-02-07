package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	app.Post("/user/:id", func(c fiber.Ctx) error {
		var user User
		if err := c.Bind().Body(&user); err != nil {
			return err
		}
		return c.JSON(user)
	})

	app.Get("/", func(c fiber.Ctx) error {

		if c.IP() == "192.168.1.1" {
			return c.Drop()
		}
		return c.SendString("Hello World")
	})

	app.Get("/sse", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")
		c.Set("Cache-Control", "max-age=0")

		return c.SendStreamWriter(func(w *bufio.Writer) {
			fmt.Fprintln(w, "event: my-message\n", 2*2)
			fmt.Fprintf(w, "data:Hello SSh\n\n")
			if err := w.Flush(); err != nil {
				log.Println("Client disconnected!")
				return
			}
		})

	})

	app.Listen(":3000")

}
