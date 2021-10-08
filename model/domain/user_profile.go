package domain

import "time"

type UserProfile struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	FirstName string
	LastName  string
	Address   string
	Age       uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
