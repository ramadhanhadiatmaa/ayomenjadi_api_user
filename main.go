package main

import (
	"am_user/models"
	"am_user/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	models.ConnectionDatabase()

	app := fiber.New()

	app.Use(cors.New())

	routes.Route(app)

	app.Listen(":8400")
}
