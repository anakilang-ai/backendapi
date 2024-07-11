package main

import (

	"github.com/anakilang-ai/backend/routes"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define a fiber handler that wraps the net/http handler
	app.Get("/WebHook", adaptor.HTTPHandlerFunc(routes.URL))

	port := ":8080"
	app.Listen(port)
}
