package domain

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/kevin07696/produce-service/generated"
	"google.golang.org/protobuf/encoding/protojson"
)

type ProductService struct {
	repo ProductWriter
}

func NewProductService(repo ProductWriter) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// WriteProduct creates a new product.
// It takes a `request` of type *generated.Product, which contains the details of the product to be created.
// Returns a StatusCode indicating the result of the operation:
//   - StatusCode.SUCCESS if the product is successfully created.
//   - StatusCode.INVALID_ARGUMENT if the request is invalid.
//   - StatusCode.ALREADY_EXISTS if a product with the same ID already exists.
//   - StatusCode.INTERNAL_ERROR if an unexpected error occurs.

func (s *ProductService) WriteProduct(ctx context.Context, request *generated.Product) (string, StatusCode) {
	for _, o := range request.MainOption.Options {
		for _, url := range o.GetImageUrls() {
			if !ValidateBlobStorage(url) {
				slog.Error("Image URL validation failed.", "image_url", url)
				return "", StatusInvalidArgument
			}
		}
	}

	options, err := protojson.Marshal(request.Options)
	if err != nil {
		slog.Error("Failed to marshall options.", "error", err)
		return "", StatusInvalidArgument
	}

	mainOption, err := protojson.Marshal(request.MainOption)
	if err != nil {
		slog.Error("Failed to marshall main option.", "error", err)
		return "", StatusInvalidArgument
	}

	attributes, err := protojson.Marshal(request.Attributes)
	if err != nil {
		slog.Error("Failed to marshall attributes.", "error", err)
		return "", StatusInvalidArgument
	}

	product := &Product{
		Name:          request.Name,
		ThumbnailUrl:  request.ThumbnailUrl,
		CategoryName:  request.CategoryName,
		Description:   request.Description,
		CentPrice:     request.CentPrice,
		AmountSold:    request.AmountSold,
		InStock:       request.InStock,
		Rating:        request.Rating,
		Options:       options,
		GalleryOption: mainOption,
		Attributes:    attributes,
	}

	status := s.repo.WriteProduct(ctx, product)
	return product.ID.String(), status
}

// UpdateProduct updates an existing product.
// It takes a `request` of type *generated.UpdateProductRequest, which contains the updated details of the product.
// Returns a StatusCode indicating the result of the operation:
//   - StatusCode.SUCCESS if the product is successfully updated.
//   - StatusCode.INVALID_ARGUMENT if the request is invalid.
//   - StatusCode.NOT_FOUND if the product to be updated does not exist.
//   - StatusCode.INTERNAL_ERROR if an unexpected error occurs.
func (s *ProductService) UpdateProduct(ctx context.Context, request *generated.UpdateProductRequest) StatusCode {
	if request.Product == nil {
		slog.Error("UpdateProduct received nil product data")
		return StatusInvalidArgument
	}

	// Validate the product ID
	if request.Id == "" {
		slog.Error("UpdateProduct received empty product ID")
		return StatusInvalidArgument
	}

	// Validate image URLs
	if request.Product.MainOption != nil {
		for _, o := range request.Product.MainOption.Options {
			for _, url := range o.GetImageUrls() {
				if !ValidateBlobStorage(url) {
					slog.Error("Image URL validation failed.", "image_url", url)
					return StatusInvalidArgument
				}
			}
		}
	}

	// Marshal the product data to JSON
	var options, mainOption, attributes []byte
	var err error

	if request.Product.Options != nil {
		options, err = protojson.Marshal(request.Product.Options)
		if err != nil {
			slog.Error("Failed to marshall options.", "error", err)
			return StatusInvalidArgument
		}
	}

	if request.Product.MainOption != nil {
		mainOption, err = protojson.Marshal(request.Product.MainOption)
		if err != nil {
			slog.Error("Failed to marshall main option.", "error", err)
			return StatusInvalidArgument
		}
	}

	if request.Product.Attributes != nil {
		attributes, err = protojson.Marshal(request.Product.Attributes)
		if err != nil {
			slog.Error("Failed to marshall attributes.", "error", err)
			return StatusInvalidArgument
		}
	}

	// Convert UUID string to UUID type
	productID, err := uuid.Parse(request.Id)
	if err != nil {
		slog.Error("Failed to parse product ID.", "error", err, "id", request.Id)
		return StatusInvalidArgument
	}

	// Create product model for repository update
	product := Product{
		Name:          request.Product.Name,
		ThumbnailUrl:  request.Product.ThumbnailUrl,
		CategoryName:  request.Product.CategoryName,
		Description:   request.Product.Description,
		CentPrice:     request.Product.CentPrice,
		AmountSold:    request.Product.AmountSold,
		InStock:       request.Product.InStock,
		Rating:        request.Product.Rating,
		Options:       options,
		GalleryOption: mainOption,
		Attributes:    attributes,
	}

	// Call repository to update the product
	return s.repo.UpdateProduct(ctx, productID, product)
}

// DeleteProduct soft deletes a product.
// It takes a `request` of type *generated.DeleteProductRequest, which contains the ID of the product to be deleted.
// Returns a StatusCode indicating the result of the operation:
//   - StatusCode.SUCCESS if the product is successfully deleted.
//   - StatusCode.INVALID_ARGUMENT if the request is invalid.
//   - StatusCode.NOT_FOUND if the product to be deleted does not exist.
//   - StatusCode.INTERNAL_ERROR if an unexpected error occurs.
func (s *ProductService) DeleteProduct(ctx context.Context, request *generated.DeleteProductRequest) StatusCode {
	// Validate the product ID
	if request.Id == "" {
		slog.Error("DeleteProduct received empty product ID")
		return StatusInvalidArgument
	}

	// Convert UUID string to UUID type
	productID, err := uuid.Parse(request.Id)
	if err != nil {
		slog.Error("Failed to parse product ID.", "error", err, "id", request.Id)
		return StatusInvalidArgument
	}

	return s.repo.DeleteProduct(ctx, productID)
}

// RecoverProduct updates a product.
// It takes a `request` of type *generated.RecoverProductRequest, which contains the ID of the product to be recovered.
// Returns a StatusCode indicating the result of the operation:
//   - StatusCode.SUCCESS if the product is successfully recovered.
//   - StatusCode.INVALID_ARGUMENT if the request is invalid.
//   - StatusCode.NOT_FOUND if the product to be recovered does not exist.
//   - StatusCode.INTERNAL_ERROR if an unexpected error occurs.
func (s *ProductService) RecoverProduct(ctx context.Context, request *generated.RecoverProductRequest) StatusCode {
	// Validate the product ID
	if request.Id == "" {
		slog.Error("RecoverProduct received empty product ID")
		return StatusInvalidArgument
	}

	// Convert UUID string to UUID type
	productID, err := uuid.Parse(request.Id)
	if err != nil {
		slog.Error("Failed to parse product ID.", "error", err, "id", request.Id)
		return StatusInvalidArgument
	}

	return s.repo.RecoverProduct(ctx, productID)
}
