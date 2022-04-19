package main

import (
	database "github.com/weldonla/OneCauseGolangLogin/databases"
	router "github.com/weldonla/OneCauseGolangLogin/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	router.Setup(app)

	app.Listen(":8000")
}
