package models

import (
	"gorm.io/gorm"
	"time"
)

type Tabler interface {
	TableName() string
}

type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Status    int            `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Category) TableName() string {
	return "categories"
}
