package models

import (
	"errors"
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

func GetAlbums() ([]Album, error) {
	var albums []Album

	result := db.Find(&albums)
	if result.Error != nil {
		return []Album{}, result.Error
	}

	return albums, nil
}

func GetAlbumById(albumId uint) (Album, error) {
	var album Album

	if err := db.First(&album, albumId).Error; err != nil {
		return album, errors.New("Album not found")
	}

	return album, nil
}

func (album *Album) CreateAlbum() (*Album, error) {
	err := db.Create(&album).Error
	if err != nil {
		return &Album{}, err
	}

	return album, nil
}
