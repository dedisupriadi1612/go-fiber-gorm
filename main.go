package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initial database postgres
	database.DatabaseInit()

	app := fiber.New()

	//initial router
	router.RouterInit(app)

	app.Listen(":3000")
}
