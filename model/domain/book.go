package domain

import (
	"time"
)

type Book struct {
	ID          uint    `gorm:"primaryKey"`
	Title       string  `gorm:"not null;unique"`
	Description string
	Price       uint    `gorm:"not null" `
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
