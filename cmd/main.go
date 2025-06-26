// @title Melibrary API
// @version 1.0
// @description A digital library API
// @host localhost:3000
// @BasePath /
// @schemes http

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/melibrary-api/config"
	"github.com/ihsankarim/melibrary-api/database"
	_ "github.com/ihsankarim/melibrary-api/docs"
	"github.com/ihsankarim/melibrary-api/routes"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
