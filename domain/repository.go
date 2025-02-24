package domain

import (
	"fmt"
	"log/slog"

	"github.com/kevin07696/produce-service/generated"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db           *gorm.DB
	errorHandler map[error]StatusCode
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{
		db: db,
		errorHandler: map[error]StatusCode{
			gorm.ErrRecordNotFound: StatusNotFound,
		},
	}
}

func (r ProductRepository) Migrate() {
	if !(r.db.Migrator().HasTable(&Product{})) {
		if err := r.db.Migrator().CreateTable(&Product{}); err != nil {
			slog.Error("Failed to migrate table products.", "error", err)
		}
	}
}

func (r ProductRepository) ReadProductSummaries(request *generated.ProductSummaryRequest) ([]ProductSummary, StatusCode) {
	query := r.db.Table(Product{}.TableName())

	// Apply stock filter if specified
	if request.Stock != generated.ProductStock_STOCK_UNSPECIFIED {
		query = query.Where("in_stock = ?", request.Stock == generated.ProductStock_STOCK_IN)
	}

	// Apply category filter if specified
	if len(request.Categories) > 0 {
		query = query.Where("category IN ?", request.Categories)
	}

	// Apply price range filter
	if request.MaxPrice > request.MinPrice {
		query = query.Where("price BETWEEN ? AND ?", request.MinPrice, request.MaxPrice)
	}

	// Determine sorting order
	var sortField string
	switch request.Sort {
	case generated.ProductSort_SORT_PRICE:
		sortField = "price"
	default:
		sortField = "created_at"
	}

	sortOrder := "ASC"
	if request.Desc {
		sortOrder = "DESC"
	}

	var summaries []ProductSummary
	result := query.Order(fmt.Sprintf("%s %s", sortField, sortOrder)).Find(&summaries)

	if result.Error != nil {
		slog.Error("Failed to read product summaries.", "error", result.Error)
		if status, exists := r.errorHandler[result.Error]; exists {
			return nil, status
		}
		return nil, StatusInternal
	}

	return summaries, StatusOK
}

func (r ProductRepository) ReadProductDetail(productID uint) (detail ProductDetail, status StatusCode) {
	var ok bool

	result := r.db.Model(&Product{}).Where(gorm.Model{ID: productID}).Take(&detail)
	if result.Error != nil {
		slog.Error("Failed to query product details", "id", productID, "error", result.Error)
		status, ok = r.errorHandler[result.Error]
		if !ok {
			status = StatusInternal
		}
	}
	return
}

func (r ProductRepository) ReadCategories() (categories []string, status StatusCode) {
	result := r.db.Table(Product{}.TableName()).Distinct("category").Find(&categories)
	if result.Error != nil {
		slog.Error("Failed to query categories", "error", result.Error)

		var ok bool
		status, ok = r.errorHandler[result.Error]
		if !ok {
			status = StatusInternal
		}
		return // Ensures we don't fall through to setting StatusOK
	}

	// Explicitly set status on success
	status = StatusOK
	return
}

func (r ProductRepository) WriteProducts() StatusCode {
	products := []Product{
		{
			Name:         "Men's T-Shirt",
			Category:     "Tops",
			Price:        1999,
			ThumbnailURL: "https://example.com/images/tshirt.jpg",
			Description:  "Cotton crew-neck t-shirt available in multiple colors.",
			ImagesURLs:   []string{"https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg"},
			Rating:       4.5,
			InStock:      true,
		},
		{
			Name:         "Women's Hoodie",
			Category:     "Tops",
			Price:        3999,
			ThumbnailURL: "https://example.com/images/hoodie.jpg",
			Description:  "Cozy fleece hoodie with front pockets.",
			ImagesURLs:   []string{"https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg"},
			Rating:       4.7,
			InStock:      true,
		},
		{
			Name:         "Leather Jacket",
			Category:     "Tops",
			Price:        12999,
			ThumbnailURL: "https://example.com/images/jacket.jpg",
			Description:  "Premium leather jacket for a stylish look.",
			ImagesURLs:   []string{"https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg"},
			Rating:       4.6,
			InStock:      false,
		},
		{
			Name:         "Running Shoes",
			Category:     "Shoes",
			Price:        7999,
			ThumbnailURL: "https://example.com/images/shoes.jpg",
			Description:  "Lightweight and comfortable running shoes.",
			ImagesURLs:   []string{"https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg"},
			Rating:       4.8,
			InStock:      true,
		},
		{
			Name:         "Casual Backpack",
			Category:     "Backpack",
			Price:        4999,
			ThumbnailURL: "https://example.com/images/backpack.jpg",
			Description:  "Durable backpack with multiple compartments.",
			ImagesURLs:   []string{"https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg"},
			Rating:       4.5,
			InStock:      true,
		},
		{
			Name:         "Leather Handbag",
			Category:     "Handbag",
			Price:        8999,
			ThumbnailURL: "https://example.com/images/handbag.jpg",
			Description:  "Elegant leather handbag for everyday use.",
			ImagesURLs:   []string{"https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg"},
			Rating:       4.7,
			InStock:      true,
		},
		{
			Name:         "Travel Duffel Bag",
			Category:     "Bag",
			Price:        6999,
			ThumbnailURL: "https://example.com/images/duffel.jpg",
			Description:  "Spacious duffel bag for travel and gym.",
			ImagesURLs:   []string{"https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg"},
			Rating:       4.6,
			InStock:      false,
		},
		{
			Name:         "Messenger Bag",
			Category:     "Bag",
			Price:        5999,
			ThumbnailURL: "https://example.com/images/messenger.jpg",
			Description:  "Stylish and practical messenger bag for work and casual use.",
			ImagesURLs:   []string{"https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg", "https://example.com/images/tshirt.jpg"},
			Rating:       4.4,
			InStock:      true,
		},
	}

	result := r.db.Create(&products)
	if result.Error != nil {
		slog.Error("Failed to write products.", "error", result.Error)
		if status, exists := r.errorHandler[result.Error]; exists {
			return status
		}
		return StatusInternal
	}

	return StatusOK
}
