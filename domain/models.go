package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"` // generate uuid is not supported on SQLite
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	Model
	Name          string         `gorm:"type:text;not null"`
	ThumbnailUrl  string         `gorm:"type:text;not null"`
	CategoryName  string         `gorm:"type:text;not null"`
	Description   string         `gorm:"type:text;not null"`
	CentPrice     uint32         `gorm:"type:uint;default:0"`
	AmountSold    uint32         `gorm:"type:uint;default:0"`
	InStock       bool           `gorm:"type:boolean;default:false"`
	Rating        float32        `gorm:"type:numeric;default:0"`
	Options       datatypes.JSON `gorm:"type:jsonb"`
	GalleryOption datatypes.JSON `gorm:"type:jsonb"`
	Attributes    datatypes.JSON `gorm:"type:jsonb"`
}

func (Product) TableName() string {
	return "products"
}

type ProductSummary struct {
	ID           string    `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	ThumbnailUrl string    `gorm:"column:thumbnail_url"`
	CategoryName string    `gorm:"column:category_name"`
	CentPrice    uint32    `gorm:"column:cent_price"`
	InStock      bool      `gorm:"column:in_stock"`
	Rating       float32   `gorm:"column:rating"`
	AmountSold   uint32    `gorm:"column:amount_sold"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

type ProductDetail struct {
	ID           string         `gorm:"column:id"`
	Name         string         `gorm:"column:name"`
	CategoryName string         `gorm:"column:category_name"`
	Description  string         `gorm:"column:description"`
	CentPrice    uint32         `gorm:"column:cent_price"`
	Rating       float32        `gorm:"column:rating"`
	Attributes   datatypes.JSON `gorm:"type:jsonb"`
	MainOption   datatypes.JSON `gorm:"type:jsonb"`
	Options      datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
}
