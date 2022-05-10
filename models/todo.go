package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	// gorm.Model

	ID              uint           `json:"id"`
	ActivityGroupId uint           `json:"activity_group_id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"delete_at"`
	Title           string         `json:"title"`
	IsActive        string         `json:"is_active"`
	Priority        string         `json:"priority"`
}
