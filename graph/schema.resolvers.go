package graph

import (
"context"
"fmt"
"strconv"

"github.com/ryanmoralesaz/colossal-stack/graph/model"
"github.com/ryanmoralesaz/colossal-stack/models"
)

// ID is the resolver for the id field.
func (r *bookResolver) ID(ctx context.Context, obj *models.Book) (string, error) {
return strconv.FormatUint(uint64(obj.ID), 10), nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *bookResolver) CreatedAt(ctx context.Context, obj *models.Book) (string, error) {
return obj.CreatedAt.Format("2006-01-02T15:04:05Z07:00"), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *bookResolver) UpdatedAt(ctx context.Context, obj *models.Book) (string, error) {
return obj.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"), nil
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.CreateBookInput) (*models.Book, error) {
book := &models.Book{
Title:  input.Title,
Author: input.Author,
}

// Handle optional publisher field
if input.Publisher != nil {
book.Publisher = *input.Publisher
}

if err := r.DB.Create(book).Error; err != nil {
return nil, fmt.Errorf("failed to create book: %w", err)
}

return book, nil
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, input model.UpdateBookInput) (*models.Book, error) {
book := &models.Book{}

// Find existing book
if err := r.DB.First(book, id).Error; err != nil {
return nil, fmt.Errorf("book not found: %w", err)
}

// Update fields if provided
if input.Title != nil {
book.Title = *input.Title
}
if input.Author != nil {
book.Author = *input.Author
}
if input.Publisher != nil {
book.Publisher = *input.Publisher
}

// Save changes
if err := r.DB.Save(book).Error; err != nil {
return nil, fmt.Errorf("failed to update book: %w", err)
}

return book, nil
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (bool, error) {
bookID, err := strconv.ParseUint(id, 10, 32)
if err != nil {
return false, fmt.Errorf("invalid ID format: %w", err)
}

if err := r.DB.Delete(&models.Book{}, bookID).Error; err != nil {
return false, fmt.Errorf("failed to delete book: %w", err)
}

return true, nil
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*models.Book, error) {
var books []*models.Book

if err := r.DB.Find(&books).Error; err != nil {
return nil, fmt.Errorf("failed to fetch books: %w", err)
}

return books, nil
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id string) (*models.Book, error) {
book := &models.Book{}

if err := r.DB.First(book, id).Error; err != nil {
return nil, fmt.Errorf("book not found: %w", err)
}

return book, nil
}

// Book returns BookResolver implementation.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
