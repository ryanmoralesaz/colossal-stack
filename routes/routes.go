package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanmoralesaz/colossal-stack/middleware"
	"github.com/ryanmoralesaz/colossal-stack/models"
)

func SetupRoutes(app *fiber.App, repo *models.Repository, authRepo *models.AuthRepository) {
	api := app.Group("/api")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authRepo.Register)
	auth.Post("/login", authRepo.Login)

	// Book routes
	books := api.Group("/books")

	// Public routes (read-only)
	books.Get("/", repo.GetBooks)
	books.Get("/:id", repo.GetBookByID)

	// Protected routes (require JWT)
	books.Post("/", middleware.Protected(), repo.CreateBook)
	books.Put("/:id", middleware.Protected(), repo.UpdateBook)
	books.Delete("/:id", middleware.Protected(), repo.DeleteBook)
}
