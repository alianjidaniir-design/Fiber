package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	// --- مسیر محافظت شده ---
	// در اینجا، "Secret Area" همان چیزی است که در زمینه HTTP به عنوان "Realm" شناخته می‌شود.
	app.Use(basicauth.New(basicauth.Config{
		Authorizer: func(user, pass string) bool {
			// منطق اعتبارسنجی
			return user == "user" && pass == "pass"
		},
		// **اینجا "Realm" تعریف می‌شود:**
		Realm: "Restricted Access Zone",
	}))

	app.Get("/protected", func(c *fiber.Ctx) error {
		return c.SendString("Access Granted to the Realm!")
	})
	log.Fatal(app.Listen(":3020"))
}
