package main

import (
	"github.com/MdHasib01/go_postgres/handler"
	"github.com/gofiber/fiber/v2"
)	

func setupRoutes(app *fiber.App) {
	app.Get("/", handler.Home)

	app.Get("/facts", handler.GetFacts)
}