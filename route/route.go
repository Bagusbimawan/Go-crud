package route

import (
	"crud-go/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Get("/user",controller.GetAllUser)
	api.Post("/user",controller.CreateUser)
	api.Post("/product",controller.CreateProduct)
	api.Get("/product",controller.GetAllProduct)
}