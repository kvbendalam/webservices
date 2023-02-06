package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kvbendalam/webservices/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
