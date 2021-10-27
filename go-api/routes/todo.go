package routes

import (
	"github.com/asamirr/Todo-Golang-Vue/controllers"
	"github.com/gofiber/fiber/v2"
)

func Todo(route fiber.Router) {
	route.Get("", controllers.GetAll)

}
