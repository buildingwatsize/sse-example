package main

import (
	"sseexample/pkg/computed"
	"sseexample/pkg/sse"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Cache-Control",
		AllowCredentials: true,
	}))

	app.Static("/", "public")

	app.Get("/events", sse.Handler)
	app.Post("/computed", computed.Handler)

	PORT := "7777"
	app.Listen(":" + PORT)
}
