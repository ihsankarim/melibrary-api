package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ihsankarim/melibrary-api/config"
	"github.com/ihsankarim/melibrary-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Input"})
	}

	var existing models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already registered"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		Role:     "user",
	}
	config.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered",
	})
}

func Login(c *fiber.Ctx) error {
	var input models.User
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON((fiber.Map{
			"error": "Invalid input",
		}))
	}

	config.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}
