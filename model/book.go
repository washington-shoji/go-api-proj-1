package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Book struct
// type Book struct {as
// 	gorm.Model
// 	ID          uuid.UUID `gorm:"type:uuid; index:unique"`
// 	Title       string    `gorm:"primary_key; index:unique" json:"title"`
// 	Subtitle    string    `json:"subtitle"`
// 	Description string    `json:"description"`
// }

type Book struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key; type:uuid"`
	Title       string    `gorm:"index:unique" json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	PublisherID uuid.UUID `json:"publisher-id"`
	Authors     []Author  `gorm:"many2many:author_books;"`
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
