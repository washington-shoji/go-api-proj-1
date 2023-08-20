package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid; index:unique"`
	Name      string    `gorm:"primary_key; index:unique" json:"name"`
	BookTitle string    `json:"book-title"`
	Books     []Book    `gorm:"foreignKey:Title; references:BookTitle"`
}

// Books struct
type Authors struct {
	Author []Author `json:"authors"`
}

// Assigning uuid
func (author *Author) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	author.ID = uuid.New()
	return
}
