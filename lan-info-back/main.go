package main

import (
	"log"
	"os"

	"0chaos.eu/lan-info-back/middlewares"
	"0chaos.eu/lan-info-back/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	main_serve()
}

func main_serve() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(etag.New(etag.Config{Weak: true}))

	routes.SetupRoutes(app)

	// Run application
	port := ":3000"
	if value, found := os.LookupEnv("PORT"); found {
		port = value
	}
	log.Fatal(app.Listen(port))
}
