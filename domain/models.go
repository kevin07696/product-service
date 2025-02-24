package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string
	Category     string `gorm:"index"`
	Price        uint32 `gorm:"index"`
	ThumbnailURL string
	Description  string
	Rating       float32   `gorm:"index"`
	InStock      bool      `gorm:"index"`
	ImagesURLs   []string  `gorm:"serializer:json"`
	CreatedAt    time.Time `gorm:"index"`
}

func (Product) TableName() string {
	return "products"
}

type ProductSummary struct {
	ID           uint
	Name         string
	Rating       float32 `gorm:"index"`
	Price        uint32
	ThumbnailURL string
}

type ProductDetail struct {
	ID          uint
	Price       uint32
	Rating      float32 `gorm:"index"`
	Name        string
	Category    string
	Description string
	ImagesURLs  []string `gorm:"serializer:json"`
	CreatedAt   time.Time
}
