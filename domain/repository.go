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

func (r *ProductRepository) handleError(ctx context.Context, message string, err error, attrs ...slog.Attr) StatusCode {
	if err == nil {
		return StatusOK
	}

	attrs = append(attrs, slog.Any("error", err))

	slog.LogAttrs(ctx, slog.LevelError, message, attrs...)

	status, ok := r.errorHandler[err]
	if !ok {
		return StatusInternalError
	}
	return status
}

// ReadProducts fetches all products unscoped.
// The `offset` parameter is treated as a page number (starting from 0).
func (r *ProductRepository) ReadProducts(ctx context.Context) (products []Product, status StatusCode) {
	// Fetch the data with the calculated offset and limit
	result := r.db.Unscoped().Find(&products)
	status = r.handleError(ctx, "Failed to read products.", result.Error)
	return
}

// ReadProductSummaries fetches all product summaries.
func (r *ProductRepository) ReadProductSummaries(ctx context.Context) ([]ProductSummary, StatusCode) {
	var summaries []ProductSummary

	// Fetch the data with the calculated offset and limit
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Select("id, name, thumbnail_url, category_name, cent_price, in_stock, rating, amount_sold, created_at").
		Where("category_name <> ?", "Archived").
		Find(&summaries)

	return summaries, r.handleError(ctx, "Failed to read product summaries", result.Error)
}

// ReadProductDetail fetches detailed information for a specific product by its ID.
func (r *ProductRepository) ReadProductDetail(ctx context.Context, productID datatypes.UUID) (ProductDetail, StatusCode) {
	var detail ProductDetail
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Select("id, name, category_name, description, price, rating, attributes, options, main_option, created_at").
		Where("id = ?", productID).
		Where("category_name <> ?", "Archived").
		Take(&detail)
	return detail, r.handleError(ctx, "Failed to read product detail", result.Error, slog.String("product_id", productID.String()))
}

// ReadCategories fetches all distinct product categories.
func (r *ProductRepository) ReadCategories(ctx context.Context) ([]string, StatusCode) {
	var categories []string

	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Distinct("category_name").
		Where("category_name <> ?", "Archived").
		Pluck("category_name", &categories)

	return categories, r.handleError(ctx, "Failed to read categories", result.Error)
}

// WriteProduct creates a new product in the database.
func (r *ProductRepository) WriteProduct(ctx context.Context, product *Product) StatusCode {
	product.ID = datatypes.UUID(uuid.New())
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).Create(product)
	return r.handleError(ctx, "Failed to write to product", result.Error)
}

// UpdateProduct updates multiple fields for an existing product by its ID.
// Zero value fields are not added to the mutation.
func (r *ProductRepository) UpdateProduct(ctx context.Context, productID datatypes.UUID, product Product) StatusCode {
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Where("id = ?", productID).
		Updates(product)
	return r.handleError(ctx, "Failed to write product.", result.Error, slog.String("product_id", productID.String()))
}

// DeleteProduct soft deletes a product by its ID.
func (r *ProductRepository) DeleteProduct(ctx context.Context, productID datatypes.UUID) StatusCode {
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Where("id = ?", productID).
		Delete(&Product{})
	return r.handleError(ctx, "Failed to delete product.", result.Error, slog.String("product_id", productID.String()))
}

// RecoverProduct updates deleted_at on a product by its ID.
func (r *ProductRepository) RecoverProduct(ctx context.Context, productID datatypes.UUID) StatusCode {
	result := r.db.Table(Product{}.TableName()).WithContext(ctx).
		Unscoped().
		Where("id = ?", productID).
		Update("deleted_at", nil)
	return r.handleError(ctx, "Failed to recover product.", result.Error, slog.String("product_id", productID.String()))
}
