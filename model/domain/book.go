package domain

import (
	"time"
)

type Book struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Price       uint           `gorm:"type:int;not null" json:"price"`
	Rating      float32        `gorm:"type:int" json:"rating"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	//DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
