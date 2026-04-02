package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanmoralesaz/colossal-stack/models"
)

func SetupRoutes(app *fiber.App, repo *models.Repository) {
	api := app.Group("/api")

	// Book routes
	books := api.Group("/books")
	books.Post("/", repo.CreateBook)
	books.Get("/", repo.GetBooks)
	books.Get("/:id", repo.GetBookByID)
	books.Put("/:id", repo.UpdateBook)
	books.Delete("/:id", repo.DeleteBook)
}
