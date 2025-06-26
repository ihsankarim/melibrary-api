package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/melibrary-api/config"
	"github.com/ihsankarim/melibrary-api/models"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	if err := config.DB.Preload("Category").Find(&books).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get books"})
	}
	return c.JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	var input models.Book
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := config.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create book"})
	}

	return c.Status(201).JSON(input)
}
