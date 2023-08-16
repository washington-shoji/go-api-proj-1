package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Publisher struct
type Publisher struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;"`
	Name string    `json:"name"`
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
