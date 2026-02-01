package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {

	app := fiber.New()
	app.Get("/", greet)
	app.Get("/greet", greetWithHTTPReq)
	http.ListenAndServe(":3000", adaptor.FiberApp(app))

}

func greet(c *fiber.Ctx) error {
	return c.SendString("Hello World 1323")

}

func greetWithHTTPReq(c *fiber.Ctx) error {
	httpReq, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return err
	}
	return c.SendString("Request URL: " + httpReq.URL.String())
}
