package handlers

import (
	"github.com/kevin07696/produce-service/domain"
	"github.com/kevin07696/produce-service/generated"
)

type Reader interface {
	ReadProductSummaries(request *generated.ProductSummaryRequest) ([]domain.ProductSummary, domain.StatusCode)
	ReadCategories() ([]string, domain.StatusCode)
}

type ProductServicer interface {
	GetProductDetail(request *generated.ProductDetailRequest) (domain.ProductDetail, domain.StatusCode)
}
