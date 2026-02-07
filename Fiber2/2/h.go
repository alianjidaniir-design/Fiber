package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Hooks().OnPreStartupMessage(func(sm *fiber.PreStartupMessageData) error {
		sm.BannerHeader = "ALI " + sm.Version + "\n-------"
		sm.AddInfo("git-hash", "Git hash", os.Getenv("GIT_HASH"))
		sm.AddInfo("preFork", "Prefork", fmt.Sprintf("12+%v", sm.Prefork), 15)
		return nil
	})
	app.Hooks().OnPostStartupMessage(func(sm *fiber.PostStartupMessageData) error {
		if !sm.Disabled && !sm.IsChild && !sm.Prevented {
			fmt.Println("competed startup")
		}
		return nil
	})

	app.Listen(":5003")
}
