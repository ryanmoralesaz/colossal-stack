package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryanmoralesaz/colossal-stack/utils"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// CreateBook creates a new book
func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Validate book
	if err := utils.ValidateStruct(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	if err := r.DB.Create(book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create book",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created successfully",
		"book":    book,
	})
}

// UpdateBook updates an existing book
func (r *Repository) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(Book)

	if err := r.DB.First(book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	input := new(Book)
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

	// Update fields
	if input.Title != "" {
		book.Title = input.Title
	}
	if input.Author != "" {
		book.Author = input.Author
	}
	if input.Publisher != "" {
		book.Publisher = input.Publisher
	}
	if input.Price > 0 {
		book.Price = input.Price
	}
	if input.Currency != "" {
		book.Currency = input.Currency
	}

	if err := r.DB.Save(book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update book",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Book updated successfully",
		"book":    book,
	})
}

// GetBooks retrieves all books
func (r *Repository) GetBooks(c *fiber.Ctx) error {
	var books []Book

	if err := r.DB.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve books",
		})
	}

	return c.JSON(books)
}

// GetBookByID retrieves a single book
func (r *Repository) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(Book)

	if err := r.DB.First(book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.JSON(book)
}

// DeleteBook deletes a book
func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := r.DB.Delete(&Book{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete book",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
