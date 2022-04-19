package router

import (
	"github.com/weldonla/OneCauseGolangLogin/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/login", controllers.Login)
}
