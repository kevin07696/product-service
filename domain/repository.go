package domain

import (
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	ProductTableName = "products"
)

type ProductRepository struct {
	db           *gorm.DB
	chunk        int // Number of records per page
	errorHandler map[error]StatusCode
}

func NewProductRepository(db *gorm.DB, chunk int) *ProductRepository {
	return &ProductRepository{
		db:    db,
		chunk: chunk,
		errorHandler: map[error]StatusCode{
			gorm.ErrRecordNotFound: StatusNotFound,
			gorm.ErrDuplicatedKey:  StatusDuplicateKey,
		},
	}
}

func (r *ProductRepository) Migrate() error {
	if !r.db.Migrator().HasTable(&Product{}) {
		if err := r.db.Migrator().CreateTable(&Product{}); err != nil {
			slog.Error("Failed to migrate table: products.", "error", err)
			return err
		}
	}
	return nil
}

func (r *ProductRepository) handleError(err error) StatusCode {
	if err == nil {
		return StatusOK
	}

	slog.Error("Database operation failed", "error", err)

	status, ok := r.errorHandler[err]
	if !ok {
		return StatusInternal
	}
	return status
}

// ReadProductSummaries fetches a paginated list of product summaries.
// The `offset` parameter is treated as a page number (starting from 0).
func (r *ProductRepository) ReadProductSummaries(page int) ([]ProductSummary, StatusCode) {
	var summaries []ProductSummary

	// Calculate the offset based on the page number and chunk size
	offset := page * r.chunk

	// Fetch the data with the calculated offset and limit
	result := r.db.Model(&Product{}).
		Select("id, name, thumbnail_url, category_name, cent_price, in_stock, rating, amount_sold, created_at").
		Offset(offset).
		Limit(r.chunk).
		Find(&summaries)

	return summaries, r.handleError(result.Error)
}

// ReadProductDetail fetches detailed information for a specific product by its ID.
func (r *ProductRepository) ReadProductDetail(productID datatypes.UUID) (ProductDetail, StatusCode) {
	var detail ProductDetail
	result := r.db.Model(&Product{}).
		Unscoped().
		Select("id, name, category_name, description, price, rating, attributes, options, main_option, created_at").
		Where("id = ?", productID).
		Take(&detail)
	return detail, r.handleError(result.Error)
}

// ReadCategories fetches a paginated list of distinct product categories.
// The `offset` parameter is treated as a page number (starting from 0).
func (r *ProductRepository) ReadCategories(page int) ([]string, StatusCode) {
	var categories []string

	// Calculate the actual offset based on the page number and chunk size
	offset := page * r.chunk

	result := r.db.Model(&Product{}).
		Unscoped().
		Distinct("category_name").
		Limit(r.chunk).
		Offset(offset).
		Pluck("category_name", &categories)

	return categories, r.handleError(result.Error)
}

// WriteProduct creates a new product in the database.
func (r *ProductRepository) WriteProduct(product *Product) StatusCode {
	product.ID = datatypes.UUID(uuid.New())
	result := r.db.Model(&Product{}).Create(product)
	return r.handleError(result.Error)
}

// UpdateProduct updates an existing product by its ID.
func (r *ProductRepository) UpdateProduct(productID datatypes.UUID, product Product) StatusCode {
	result := r.db.Model(&Product{}).
		Where("id = ?", productID).
		Updates(product)
	return r.handleError(result.Error)
}

// DeleteProduct deletes a product by its ID.
func (r *ProductRepository) DeleteProduct(productID datatypes.UUID) StatusCode {
	result := r.db.Model(&Product{}).
		Where("id = ?", productID).
		Delete(&Product{})
	return r.handleError(result.Error)
}
