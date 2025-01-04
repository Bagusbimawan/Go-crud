package main

import (
	
	"crud-go/database"
	"crud-go/route"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)
var DB *gorm.DB
func main() {
	app := fiber.New()
	database.ConnectDatabase()
	
	route.SetupRoutes(app)
	app.Listen(":3000")

}
