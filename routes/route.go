package routes

import (
	"am_user/controllers"
	"am_user/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	user := r.Group("/api")

	user.Get("/", middlewares.Auth, controllers.Index)
	user.Get("/:email", middlewares.Auth, controllers.Show)
	user.Get("/:username", middlewares.Auth, controllers.Check)
	user.Post("/", middlewares.Auth, controllers.Create)
	user.Put("/:username", middlewares.Auth, controllers.Update)
	user.Delete("/:username", middlewares.Auth, controllers.Delete)
}
