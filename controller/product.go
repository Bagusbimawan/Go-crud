package controller

import (
	"crud-go/database"
	"crud-go/model"

	"github.com/gofiber/fiber/v2"
)


func CreateProduct(c *fiber.Ctx)error{
	var Products model.Product
	if err:= c.BodyParser(&Products); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"status": 400,
		})
	}

	if err:= database.DB.Create(&Products).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create product",
			"status": 500,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product created successfully",
		"status": 200,
		"data": Products,
	})
}

func GetAllProduct(C *fiber.Ctx)error{
	var products []model.Product

	if err := database.DB.Find(&products).Error ; err != nil {
		return C.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all product",
			"status": 500,
		})
	}

	return C.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Product get Success",
		"status":200,
		"data": products,
	})
}