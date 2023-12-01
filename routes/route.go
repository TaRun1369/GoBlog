package routes

import (
	"github.com/TaRun1369/GoBlog/controller"
	"github.com/TaRun1369/GoBlog/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register) // Register is the function in controller
	app.Use(middleware.IsAuthenticate)
}
