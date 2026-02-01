package main

import (
	"fmt"
	"io"
	"log"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.BaseURL())
		fmt.Println(c.Get("Referer"))
		return c.SendString("Hello, World!")
	})

	req := httptest.NewRequest("GET", "http://google.com", nil)
	req.Header.Set("Referer", "good afternoon")

	resp, _ := app.Test(req)

	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

	log.Fatal(app.Listen(":3000"))

}
