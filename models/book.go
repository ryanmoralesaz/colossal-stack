package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model // embedded struct
	// ID uint // primary key
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
	Title     string  `json:"title" gorm:"not null"` // metadata and integrity schema
	Author    string  `json:"author" gorm:not null"`
	Publisher string  `json:"publisher"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2)" validate:"gte=0"`
	Currency  string  `json:"currency" gorm:"size:3;default:'USD'" validate:"len=3"`
}

func MigrateBooks(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
