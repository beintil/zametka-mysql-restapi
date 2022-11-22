package main

import (
	"github.com/beintil/zametka-mysql-restapi/database"
	"github.com/gofiber/fiber/v2"
	"github.com/beintil/zametka-mysql-restapi/controller"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()

	app.Get("/api/zametka", controllers.AllZametka)
	app.Post("/api/zametka", controllers.Create)
	app.Delete("/api/zametka/:id", controllers.Delete)

	app.Listen(":8080")
}