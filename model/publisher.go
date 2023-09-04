package model

import (
	//"github.com/google/uuid"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Publisher struct
// type Publisher struct {
// 	gorm.Model
// 	ID        uuid.UUID `gorm:"primary_key; type:uuid; index:unique"`
// 	Name      string    `gorm:"primary_key; index:unique" json:"name"`
// 	BookTitle string    `json:"book-title"`
// 	Books     []Book    `gorm:"foreignKey:Title; references:BookTitle"`
// }

type Publisher struct {
	gorm.Model
	ID    uuid.UUID `gorm:"primary_key; type:uuid"`
	Name  string    `gorm:"index:unique" json:"name"`
	Books []Book    `gorm:"foreignKey:PublisherID"`
}

// Publishers struct
type Publishers struct {
	Publishers []Publisher `json:"publishers"`
}

// Assigning uuid
func (publisher *Publisher) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	publisher.ID = uuid.New()
	return
}
