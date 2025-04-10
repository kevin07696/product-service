package domain

import (
	"context"

	"github.com/google/uuid"
)

type ProductWriter interface {
	WriteProduct(ctx context.Context, product *Product) StatusCode
	UpdateProduct(ctx context.Context, productID uuid.UUID, product Product) StatusCode
	DeleteProduct(ctx context.Context, productID uuid.UUID) StatusCode
	RecoverProduct(ctx context.Context, productID uuid.UUID) StatusCode
}
