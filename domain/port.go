package domain

import (
	"context"

	"gorm.io/datatypes"
)

type ProductWriter interface {
	WriteProduct(ctx context.Context, product *Product) StatusCode
	UpdateProduct(ctx context.Context, productID datatypes.UUID, product Product) StatusCode
	DeleteProduct(ctx context.Context, productID datatypes.UUID) StatusCode
	RecoverProduct(ctx context.Context, productID datatypes.UUID) StatusCode
}
