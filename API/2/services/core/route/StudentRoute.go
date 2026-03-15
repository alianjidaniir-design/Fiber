package route

import (
	"github.com/gofiber/fiber/v2"
)

var StudentRoute = map[string]string{
	"StudentCreate": "student/create",
}

func SetupStudentRoute(app *fiber.App) map[string]string {
	app.Post("StudentCreate", Create)
	return StudentRoute
}
