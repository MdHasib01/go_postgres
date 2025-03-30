package main

import (
	"github.com/MdHasib01/go_postgres/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	

	app.Listen(":8000")
}