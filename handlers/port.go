package handlers

import (
	"github.com/kevin07696/produce-service/domain"
	"gorm.io/datatypes"
)

type Reader interface {
	ReadProductSummaries() ([]domain.ProductSummary, domain.StatusCode)
	ReadProductDetail(productID datatypes.UUID) (domain.ProductDetail, domain.StatusCode)
	ReadCategories() ([]string, domain.StatusCode)
}
