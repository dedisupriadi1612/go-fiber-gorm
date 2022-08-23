package router

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"Hello": "World",
		})
	})
}
