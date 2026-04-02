package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"-" gorm:"not null"`
	RefreshToken string `json:"-" gorm:"type:text"`
	IsAdmin      bool   `json:"is_admin" gorm:"default:false"`
}

// HashPassword hashes the user's password
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies password
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// MigrateUsers runs auto-migration for User table
func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
