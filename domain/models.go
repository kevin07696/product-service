package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string
	Category     string  `gorm:"index"`
	Price        float32 `gorm:"index"`
	ThumbnailURL string
	ImagesURLs   []string `gorm:"type:text"` 
	Description  string
	Rating       float32   `gorm:"index"`
	InStock      bool      `gorm:"index"`
	CreatedAt    time.Time `gorm:"index"`
}

func (Product) TableName() string {
	return "products"
}

type ProductSummary struct {
	ID           uint
	Name         string
	Price        float32
	ThumbnailURL string
}

type ProductDetail struct {
	ID          uint
	Name        string
	Category    string
	Price       float32
	ImagesURLs  []string
	Description string
	Stock       uint
	CreatedAt   time.Time
}
