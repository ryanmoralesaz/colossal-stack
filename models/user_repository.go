package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanmoralesaz/colossal-stack/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type CreateUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type DeleteUserInput struct {
	Email string `json:"email" validate:"required,email"`
}

// GetUsers returns all users (admin only - requires auth)
func (r *UserRepository) GetUsers(c *fiber.Ctx) error {
	var users []User

	if err := r.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return c.JSON(users)
}

// CreateUser creates a new user (admin only - requires auth)
func (r *UserRepository) CreateUser(c *fiber.Ctx) error {
	input := new(CreateUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Validate input
	if err := utils.ValidateStruct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	// Check if user exists
	var existingUser User
	if err := r.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	// Create user
	user := User{
		Email: input.Email,
	}

	if err := user.HashPassword(input.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	if err := r.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

// DeleteUser deletes a user by email (admin only - requires auth)
func (r *UserRepository) DeleteUser(c *fiber.Ctx) error {
	email := c.Params("email")

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email parameter required",
		})
	}

	// Find user
	var user User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Delete user
	if err := r.DB.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
		"email":   email,
	})
}

// GetCurrentUser returns the authenticated user's info
func (r *UserRepository) GetCurrentUser(c *fiber.Ctx) error {
	// Get user ID from auth middleware context
	userID := c.Locals("userID").(uint)

	var user User
	if err := r.DB.Select("id", "email", "created_at", "updated_at").First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}
