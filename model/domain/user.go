package domain

import "time"

type User struct {
	ID          uint        `gorm:"primaryKey"`
	Email       string      `gorm:"not null;unique"`
	Password    string      `gorm:"not null"`
	UserProfile UserProfile `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Books       []Book      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
