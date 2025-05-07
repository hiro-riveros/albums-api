package models

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey"`
	Title     string         `json:"title"`
	Artist    string         `json:"artist"`
	Price     float64        `json:"price"`
	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
