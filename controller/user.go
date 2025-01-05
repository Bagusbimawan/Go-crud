package controller

import (
	"crud-go/database"
	"crud-go/model"

	"github.com/gofiber/fiber/v2"
)


func CreateUser(C *fiber.Ctx) error {
	var user model.User
	if err := C.BodyParser(&user); err != nil {
		C.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "failed",
		})
		return err
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return C.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "failed",
		})
	}

	return C.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
		"status":  "success",
		"user":    user,
	})
}

func GetAllUser(C *fiber.Ctx) error {
	var users []model.User
	if err:= database.DB.Find(&users).Error; err != nil{
		return C.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Failed to get all user",
			"status":500,
		})
	}

	return C.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Success Get All user",
		"status":200,
		"data":users,
	})
}