package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/melibrary-api/controllers"
	"github.com/ihsankarim/melibrary-api/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)

	api := app.Group("/api", middleware.JWTProtected())
	api.Get("/books", controllers.GetBooks)
	api.Post("/books", controllers.CreateBook)
	api.Get("/category", controllers.GetCategories)
	api.Post("/category", controllers.CreateCategory)
}
