package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanmoralesaz/colossal-stack/middleware"
	"github.com/ryanmoralesaz/colossal-stack/models"
)

func SetupRoutes(app *fiber.App, repo *models.Repository, authRepo *models.AuthRepository, userRepo *models.UserRepository) {
	api := app.Group("/api")

	// Auth routes (public) with rate limiting
	auth := api.Group("/auth")
	auth.Post("/register", middleware.AuthRateLimit(), authRepo.Register)
	auth.Post("/login", middleware.AuthRateLimit(), authRepo.Login)
	auth.Post("/refresh", authRepo.RefreshToken)

	// User routes (protected)
	users := api.Group("/users", middleware.Protected())
	users.Get("/", userRepo.GetUsers)              // List all users
	users.Get("/me", userRepo.GetCurrentUser)      // Get current user
	users.Post("/", userRepo.CreateUser)           // Create new user
	users.Delete("/:email", userRepo.DeleteUser)   // Delete user by email

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
