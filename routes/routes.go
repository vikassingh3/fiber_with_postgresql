package routes

import (
	"project/controllers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	app.Get("/", controllers.GetAllbooks)
	app.Get("/:id", controllers.Getbook)
	app.Get("/", controllers.Createbook)
	app.Delete("/:id", controllers.Deletebook)
}
