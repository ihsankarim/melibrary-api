package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/melibrary-api/config"
	"github.com/ihsankarim/melibrary-api/database"
	"github.com/ihsankarim/melibrary-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(
			"Error loading .env file",
		)
	}

	config.InitDatabase()
	database.Migrate()

	app := fiber.New()
	routes.Setup(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
