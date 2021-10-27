package routes

import (
	"github.com/asamirr/Todo-Golang-Vue/controllers"
	"github.com/gofiber/fiber/v2"
)

func Todo(route fiber.Router) {
	route.Get("", controllers.GetAll)
	route.Get("/:id", controllers.GetTask)
	route.Post("", controllers.PostTask)
	route.Put("/:id", controllers.UpdateTask)
	route.Delete("/:id", controllers.DeleteTask)

}
