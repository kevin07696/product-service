package handlers

import (
	"context"

	"github.com/kevin07696/produce-service/domain"
	"github.com/kevin07696/produce-service/generated"
	"gorm.io/datatypes"
)

type Reader interface {
	ReadProductSummaries(context.Context) ([]domain.ProductSummary, domain.StatusCode)
	ReadProductDetail(context.Context, datatypes.UUID) (domain.ProductDetail, domain.StatusCode)
	ReadCategories(context.Context) ([]string, domain.StatusCode)
}

type Service interface {
	WriteProduct(ctx context.Context, request *generated.Product) (string, domain.StatusCode)
	UpdateProduct(ctx context.Context, request *generated.UpdateProductRequest) domain.StatusCode
	DeleteProduct(ctx context.Context, request *generated.DeleteProductRequest) domain.StatusCode
	RecoverProduct(ctx context.Context, request *generated.RecoverProductRequest) domain.StatusCode
}
