package main

import (
	"log"

	"github.com/asamirr/Todo-Golang-Vue/config"
	"github.com/asamirr/Todo-Golang-Vue/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 		"message": "hello main",
	// 	})
	// })

	api := app.Group("/api")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "api endpoint",
		})
	})

	routes.Todo(api.Group("/tasks"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	setupRoutes(app)

	err = app.Listen(":8000")

	if err != nil {
		panic(err)
	}
}
