package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// handle POST /api/boos with CreateBook
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	// book = {Title: "", "Author: "", Publisher: ""
	if err := context.BodyParser(&book); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not parse request body",
			"error":   err.Error(),
		})
	}

	if err := r.DB.Create(&book).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{

			"message": "Could not create book",
			"error":   err.Error(),
		})
	}
	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created succesfully",
		"data":    book,
	})

}

// handle GET route /api/books
func (r *Repository) GetBooks(context *fiber.Ctx) error {
	books := []Book{}

	if err := r.DB.Find(&books).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Books fetched successfully",
			"data":    books,
		})
	}
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Books fetched successfully",
		"data": books,
	})
}

// handle GET book by id /api/books/:id
func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	book := Book{}

	if id == "" {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "ID cannot be empty",
		})
	}

	if err := r.DB.Where("id=?", id).First(&book).Error; err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
			"error":   err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book fetched successfully",
		"data":    book,
	})
}

// handle update book PUT /api/books/:id
func (r *Repository) UpdateBook(context *fiber.Ctx) error {
	id := context.Params("id")
	book := Book{}

	if err := r.DB.Where("id=?",id).First(&book).Error; err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}
	updateData := Book {}
	if err := context.BodyParser(&updateData); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not parse request body",
		})
	}
	book.Title = updateData.Title
	book.Author = updateData.Author
	book.Publisher = updateData.Publisher

	if err := r.DB.Save(&book).Error; err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Could not update book",
			})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book updated successfully",
		"data":    book,
	})
}

// handle book delete route DELETE /api/books/:id
func (r *Repository) DeleteBook(context *fiber.Ctx) error  {
		id := context.Params("id")
		book := Book {}

		if err := r.DB.Where("id = ?", id).Delete(&book).Error; err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not delete boook",
			})
		}
		
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Book deleted successfully",
		})
}
