package router

import (
	"go-fiber-gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func RouterInit(app *fiber.App) {
	app.Get("/", handler.App)
	app.Get("/user", handler.UserHandlerRead)
}
