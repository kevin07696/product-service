package domain

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db           *gorm.DB
	errorHandler map[error]StatusCode
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
		errorHandler: map[error]StatusCode{
			gorm.ErrRecordNotFound: StatusNotFound,
			gorm.ErrDuplicatedKey:  StatusAlreadyExists,
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
		return StatusInternalError
	}
	return status
}

// ReadProductSummaries fetches a paginated list of product summaries.
// The `offset` parameter is treated as a page number (starting from 0).
func (r *ProductRepository) ReadProductSummaries(ctx context.Context) ([]ProductSummary, StatusCode) {
	var summaries []ProductSummary

	// Fetch the data with the calculated offset and limit
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Select("id, name, thumbnail_url, category_name, cent_price, in_stock, rating, amount_sold, created_at").
		Find(&summaries)

	return summaries, r.handleError(result.Error)
}

// ReadProductDetail fetches detailed information for a specific product by its ID.
func (r *ProductRepository) ReadProductDetail(ctx context.Context, productID datatypes.UUID) (ProductDetail, StatusCode) {
	var detail ProductDetail
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Select("id, name, category_name, description, price, rating, attributes, options, main_option, created_at").
		Where("id = ?", productID).
		Take(&detail)
	return detail, r.handleError(result.Error)
}

// ReadCategories fetches a paginated list of distinct product categories.
// The `offset` parameter is treated as a page number (starting from 0).
func (r *ProductRepository) ReadCategories(ctx context.Context) ([]string, StatusCode) {
	var categories []string

	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Distinct("category_name").
		Pluck("category_name", &categories)

	return categories, r.handleError(result.Error)
}

// WriteProduct creates a new product in the database.
func (r *ProductRepository) WriteProduct(ctx context.Context, product *Product) StatusCode {
	product.ID = datatypes.UUID(uuid.New())
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).Create(product)
	return r.handleError(result.Error)
}

// UpdateProduct updates multiple fields for an existing product by its ID.
// Zero value fields are not added to the mutation.
func (r *ProductRepository) UpdateProduct(ctx context.Context, productID datatypes.UUID, product Product) StatusCode {
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Where("id = ?", productID).
		Updates(product)
	return r.handleError(result.Error)
}

// DeleteProduct soft deletes a product by its ID.
func (r *ProductRepository) DeleteProduct(ctx context.Context, productID datatypes.UUID) StatusCode {
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Where("id = ?", productID).
		Delete(&Product{})
	return r.handleError(result.Error)
}

// RecoverProduct updates deleted_at on a product by its ID.
func (r *ProductRepository) RecoverProduct(ctx context.Context, productID datatypes.UUID) StatusCode {
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Unscoped().
		Where("id = ?", productID).
		Update("deleted_at", nil)
	return r.handleError(result.Error)
}
