package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model // embedded struct
	// ID uint // primary key
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
	Title     string `json:"title" gorm:"not null"` // metadata and integrity schema
	Author    string `json:"author" gorm:not null"`
	Publisher string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
