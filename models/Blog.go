package models

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CategoryId uint           `json:"category_id"`
	Title      string         `json:"title"`
	Image      string         `json:"image"`
	Detail     string         `json:"detail" gorm:"type:text"`
	Views      int            `json:"views" gorm:"default:0"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Category   Category       `gorm:"foreignKey=CategoryId" json:"category"`
}

func (Blog) TableName() string {
	return "blogs"
}
