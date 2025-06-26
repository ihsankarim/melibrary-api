package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/melibrary-api/config"
	"github.com/ihsankarim/melibrary-api/models"
)

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get categories"})
	}
	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	var input models.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := config.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create category"})
	}
	return c.Status(201).JSON(input)
}
