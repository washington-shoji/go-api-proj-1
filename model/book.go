package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Book struct
type Book struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	Author      Author    `gorm:"embedded"`
}

// Books struct
type Books struct {
	Books []Book `json:"books"`
}

// Assigning uuid
func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	book.ID = uuid.New()
	return
}
