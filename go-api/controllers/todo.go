package controllers

import "github.com/gofiber/fiber/v2"

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{
		Id:        1,
		Title:     "Play some good music man",
		Completed: false,
	},
	{
		Id:        2,
		Title:     "What you want???",
		Completed: true,
	},
}

func GetAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"tasks": todos,
		},
	})
}
