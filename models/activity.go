package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	// gorm.Model

	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at"`
	Title     string         `json:"title"`
	Email     string         `json:"email"`
}
